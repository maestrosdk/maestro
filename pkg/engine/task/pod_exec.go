package task

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/afero"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// PodExec defines a command that will be executed in a running Pod.
// RestCfg - The REST configuration for the cluster.
// PodName - The pod name on which to execute the command.
// PodNamespace - Namespace of the pod
// ContainerName - optional container to execute the command in. If empty, first container is taken
// Args - The command (and args) to execute.
// In - Command input stream.
// Out - Command output stream
// Err - Command error stream
type PodExec struct {
	RestCfg       *rest.Config
	PodName       string
	PodNamespace  string
	ContainerName string
	Args          []string
	In            io.Reader
	Out           io.Writer
	Err           io.Writer
	TTY           bool
}

// Run executes a command in a pod. This is a distilled version of what `kubectl exec` (and
// also `kubectl  cp`) doing under the hood: a POST request is made to the `exec` subresource
// of the v1/pods endpoint containing the pod information and the command. Here is a good article
// describing it in detail: https://erkanerol.github.io/post/how-kubectl-exec-works/
func (pe *PodExec) Run() error {
	codec := serializer.NewCodecFactory(scheme.Scheme)
	restClient, err := apiutil.RESTClientForGVK(
		schema.GroupVersionKind{
			Version: "v1",
			Kind:    "pods",
		},
		pe.RestCfg,
		codec)
	if err != nil {
		return err
	}

	req := restClient.
		Post().
		Resource("pods").
		Name(pe.PodName).
		Namespace(pe.PodNamespace).
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Stdin:     pe.In != nil,
		Stdout:    pe.Out != nil,
		Stderr:    pe.Err != nil,
		TTY:       pe.TTY,
		Container: pe.ContainerName,
		Command:   pe.Args,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(pe.RestCfg, "POST", req.URL())
	if err != nil {
		return err
	}

	so := remotecommand.StreamOptions{
		Stdin:  pe.In,
		Stdout: pe.Out,
		Stderr: pe.Err,
		Tty:    pe.TTY,
	}

	// The result of the executor.Stream() call itself is returned through the streams (In, Out and Err)
	// defined in the PodExec, e.g. when downloading a file, pe.Out will  contain the file bytes.
	return exec.Stream(so)
}

// FileSize fetches the size of a file in a remote pod. It runs `stat -c %s file` command in the
// pod and parses the output.
func FileSize(file string, pod *v1.Pod, restCfg *rest.Config) (int64, error) {
	readout, stdout := io.Pipe()

	pe := PodExec{
		RestCfg:       restCfg,
		PodName:       pod.Name,
		PodNamespace:  pod.Namespace,
		ContainerName: pipePodContainerName,
		Args:          []string{"stat", "-c", "%s", file},
		In:            nil,
		Out:           stdout,
		Err:           nil,
		TTY:           true, // this will forward 2>&1. otherwise, reading from Out will never return for e.g. missing files
	}

	errCh := make(chan error, 1)
	go func() {
		defer stdout.Close()
		errCh <- pe.Run()
	}()

	buf, err := ioutil.ReadAll(readout)
	if err != nil {
		return 0, fmt.Errorf("failed to get the size of %s `: %v", file, err)
	}

	raw := string(buf)

	// THIS CHECK HAS TO HAPPEN AFTER WE CONSUME THE STDOUT/STDERR READER.
	if err := <-errCh; err != nil {
		return 0, fmt.Errorf("failed to get the size of %s, err: %v, %s", file, err, raw)
	}

	raw = raw[:len(raw)-2] // remove trailing \n\r
	size, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the size '%s' of the file %s : %v", raw, file, err)
	}

	return size, nil
}

// DownloadFile fetches a file from a remote pod. It runs `tar cf - file` command and streams contents
// of the file via the stdout. Locally, the tar file is extracted into the passed afero filesystem where
// it is saved under the same path. Afero filesystem is used to allow the caller downloading and persisting
// of multiple files concurrently (afero filesystem is thread-safe).
func DownloadFile(fs afero.Fs, file string, pod *v1.Pod, restCfg *rest.Config) error {
	readout, writeout := io.Pipe()
	stderr := bytes.Buffer{}

	pe := PodExec{
		RestCfg:       restCfg,
		PodName:       pod.Name,
		PodNamespace:  pod.Namespace,
		ContainerName: pipePodContainerName,
		Args:          []string{"tar", "cf", "-", file},
		In:            nil,
		Out:           writeout,
		Err:           &stderr,
	}

	// When downloading files using remotecommand.Executor, we pump the contents of the tar file through
	// the stdout. PodExec.Run() calls the stream executor, which first writes stdout and stderr of the
	// command and then returns with an exit code. Since we're using io.Pipe reader and writer which is
	// tread-safe but SYNCHRONOUS, Run() call will not return until we've consumed the streams (and will
	// block the execution). (for more details on how the underlying streams are copied see remotecommand/v4.go:54
	//
	// TL;DR:
	//  - execute PodExec.Run() in a goroutine
	//  - when using io.Pipe for In or Out streams, they have to be consumed first because io.Pipe is
	//    synchronous (and thread-safe)
	//  - there seems to be a bug when using BOTH Out AND Err pipe-based streams. When trying to consume
	//    both (in goroutines), one of them would end up blocking the execution ¯\_(ツ)_/¯
	//
	// See `kubectl cp` copyFromPod method for another example:
	// https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/cp/cp.go#L291
	errCh := make(chan error, 0)
	go func() {
		defer writeout.Close()
		errCh <- pe.Run()
	}()

	if err := untarFile(fs, readout, file); err != nil {
		return fmt.Errorf("failed to untar pipe file: %v", err)
	}

	// THIS CHECK HAS TO HAPPEN AFTER WE CONSUME THE STDOUT/STDERR READER.
	if err := <-errCh; err != nil {
		return fmt.Errorf("failed to copy pipe file. err: %v, stderr: %s", err, stderr.String())
	}

	return nil
}

// untarFile extracts a tar file from the passed reader using passed file name.
func untarFile(fs afero.Fs, r io.Reader, fileName string) error {
	tr := tar.NewReader(r)

	for {
		header, err := tr.Next()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		// the target location of the file. tar strips the leading "/" however, we treat the pipe file path
		// as a key to the underlying data (otherwise we'll have to start splitting paths). To avoid all
		// the complexity and because we only extract one file here, the path is taken from the PipeFile configuration
		target := fileName

		// check the file type
		switch header.Typeflag {
		// if it's a file create it
		case tar.TypeReg:
			f, err := fs.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; deferring would cause each file close
			// to wait until all operations have completed.
			f.Close() // nolint

		default:
			fmt.Printf("skipping %s because it is not a regular file or a directory", header.Name)
		}
	}

	return nil
}
