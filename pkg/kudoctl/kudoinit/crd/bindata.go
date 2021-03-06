// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package crd generated by go-bindata.// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
package crd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x3a\x6b\x73\xdb\xb6\x96\xdf\xfd\x2b\xce\xa8\x3b\x63\x3b\x95\xa8\xd8\xed\x74\x5b\xcd\x66\x33\x59\x27\xe9\x78\x9b\x87\x27\x76\xba\xd3\xb5\xdd\x5b\x88\x3c\x12\x51\x93\x00\x83\x87\x64\xb5\xe9\x7f\xbf\x73\x0e\x40\x8a\xb2\xa9\x47\xdc\xa6\xf7\xc3\xe5\x17\x5b\x20\x70\xde\x6f\x70\x6f\x30\x18\xec\x89\x4a\xfe\x88\xc6\x4a\xad\x46\x20\x2a\x89\xb7\x0e\x15\xfd\xb2\xc9\xcd\xb7\x36\x91\x7a\x38\x3b\xda\xbb\x91\x2a\x1b\xc1\x89\xb7\x4e\x97\xef\xd0\x6a\x6f\x52\x7c\x8e\x13\xa9\xa4\x93\x5a\xed\x95\xe8\x44\x26\x9c\x18\xed\x01\x08\xa5\xb4\x13\xb4\x6c\xe9\x27\x40\xaa\x95\x33\xba\x28\xd0\x0c\xa6\xa8\x92\x1b\x3f\xc6\xb1\x97\x45\x86\x86\x81\xd7\xa8\x67\x8f\x93\xaf\x93\xa3\x3d\x80\xd4\x20\x1f\xbf\x90\x25\x5a\x27\xca\x6a\x04\xca\x17\xc5\x1e\x80\x12\x25\x8e\x40\x2a\xeb\x84\x4a\xd1\x26\x37\x3e\xd3\x49\x86\xb3\x3d\x5b\x61\x4a\xc8\xa6\x46\xfb\x6a\x04\xcd\x7a\x38\x12\xe9\x08\x3c\x9c\xc6\xd3\xbc\x54\x48\xeb\x7e\x58\x59\x7e\x25\xad\xe3\x57\x55\xe1\x8d\x28\x5a\xd8\x78\xd5\x4a\x35\xf5\x85\x30\xcb\xf5\x3d\x00\x9b\xea\x0a\x47\xf0\x86\x50\x55\x22\xc5\x6c\x0f\x20\xb2\xc5\xa8\x07\x91\xf0\xd9\xd1\x18\x9d\x38\x0a\x80\xd2\x1c\x4b\x11\x08\x03\xd0\x15\xaa\x67\x67\xa7\x3f\x7e\x75\xbe\xb2\x0c\x90\xa1\x4d\x8d\xac\x1c\x4b\xa8\xa6\x11\xa4\x05\x97\x23\x84\xcd\x30\xd1\x86\x7f\x36\x94\xc2\xb3\xb3\xd3\xa4\x01\x51\x19\x5d\xa1\x71\xb2\x16\x43\x78\x5a\x4a\x6f\xad\xde\x41\xb8\x4f\x34\x85\x5d\x90\x91\xb6\x31\x20\x8e\xcc\x61\x16\xd9\x00\x3d\x01\x97\x4b\x0b\x06\x2b\x83\x16\x55\xd0\x3f\x2d\x0b\x05\x7a\xfc\x2b\xa6\x2e\x81\x73\x34\x74\x10\x6c\xae\x7d\x91\x91\x59\xcc\xd0\x38\x30\x98\xea\xa9\x92\xbf\x35\xd0\x2c\x38\xcd\x68\x0a\xe1\xd0\x3a\x90\xca\xa1\x51\xa2\x80\x99\x28\x3c\xf6\x41\xa8\x0c\x4a\xb1\x00\x83\x04\x17\xbc\x6a\x41\xe0\x2d\x36\x81\xd7\xda\x90\x40\x26\x7a\x04\xb9\x73\x95\x1d\x0d\x87\x53\xe9\x6a\x83\x4e\x75\x59\x7a\x25\xdd\x62\xc8\xb6\x29\xc7\xde\x69\x63\x87\x19\xce\xb0\x18\x5a\x39\x1d\x08\x93\xe6\xd2\x61\xea\xbc\xc1\xa1\xa8\xe4\x80\x89\x55\x6c\xd4\x49\x99\x7d\x61\xa2\x0b\xd8\xfd\x15\xe1\xb9\x05\xd9\x81\x75\x46\xaa\x69\xeb\x05\x1b\xde\x06\x29\x93\x05\x92\x4e\x45\x3c\x1a\xb8\x58\x0a\x93\x96\x48\x1e\xef\x5e\x9c\x5f\x40\x8d\x3a\x08\x3c\xc8\x76\xb9\xd5\x2e\xc5\x4c\x22\x92\x6a\x82\x26\xec\x9c\x18\x5d\x32\x14\x54\x59\xa5\xa5\x72\xfc\x23\x2d\x24\x2a\x07\xd6\x8f\x4b\xe9\x48\x7f\x1f\x3c\x5a\x47\x1a\x48\xe0\x84\x3d\x19\xc6\x08\xbe\xca\x84\xc3\x2c\x81\x53\x05\x27\xa2\xc4\xe2\x44\x58\xfc\xec\x42\x26\x69\xda\x01\x09\x6f\x37\x31\xb7\x83\xd0\xdd\xcd\x41\x4e\xad\x17\x75\xc4\x58\xa3\x93\xda\xd5\xce\x2b\x4c\x57\x4c\x3f\x43\x2b\x0d\x99\xaa\x13\x0e\xc9\xc0\xeb\x9d\xc9\x0a\xb0\x6e\xa7\x8b\xae\x6e\x84\xd3\xa6\xd3\xfb\xee\xd1\xf1\x76\x75\x37\x93\x2d\x27\x12\xc9\x58\x0c\x4e\xd0\x20\xc5\x03\xa7\xc9\x76\xc2\xab\xf4\xde\x99\xe8\x7f\xf7\x10\xad\xa7\x11\x36\x04\x88\x4e\x32\x9f\x9d\x9d\xd6\x41\x21\xc4\x02\xac\xa9\xeb\xc0\xbb\x41\x85\xf5\x33\x91\x58\x64\x67\xc2\xe5\x3b\xe0\xde\x3f\x9d\x04\x64\xec\x3a\x2c\x8a\x4a\x62\x8a\x2b\xd1\x87\x83\x23\x8a\x2c\x2e\x92\x95\x19\x8c\xef\xfa\xc1\x41\xa2\xef\x2d\xa3\x93\x13\x52\x81\x20\x67\x94\x19\xfc\xef\xf9\xdb\x37\xc3\xef\x75\xa0\x0c\x44\x9a\xa2\xb5\xc1\x08\x4a\x54\xae\x0f\xd6\xa7\x39\x08\x5b\xdb\xc7\x39\xbd\x49\x4a\xa1\xe4\x04\xad\x4b\x22\x34\x34\xf6\xf2\xf8\x3a\x81\x97\xda\x00\xde\x8a\xb2\x2a\xb0\x0f\x32\xc8\xab\xf1\xe4\x5a\xa9\xd2\x06\x66\x9a\xb3\x30\x97\x2e\x67\x92\x2a\x9d\x45\xa2\xe7\x4c\xac\x13\x37\x08\x3a\x12\xeb\x11\x0a\x79\x83\x23\xe8\x91\x45\xb4\x50\xff\x4e\x59\xe8\x8f\x1e\x1c\xcc\x73\x34\x08\x3d\xfa\xd9\x0b\x08\x9b\x90\x4b\x6b\xb5\x06\x97\x88\x5d\x2e\x1c\x38\x23\xa7\x53\x24\xdb\xe7\x28\x42\x9e\x7a\x08\xda\x10\xfd\x4a\xb7\x36\x33\x08\x92\x67\x34\xd5\xec\x1e\x21\x97\xc7\xd7\x3d\x38\x58\xe5\x0b\xa4\xca\xf0\x16\x8e\x41\xaa\xc0\x59\xa5\xb3\xc3\x04\x2e\x58\x33\x0b\xe5\xc4\x2d\xc1\x4c\x73\x6d\x51\x81\x56\xc5\x82\x28\xce\xc5\x0c\xc1\xea\x12\x61\x8e\x45\x31\x08\x7e\x9a\xc1\x5c\x2c\x88\x87\x5a\x94\xa4\x55\x01\x95\x30\xee\x4e\x42\xba\x78\xfb\xfc\xed\x28\x60\x23\xb5\x4d\x15\xa1\xa0\x90\x37\x91\x94\x6e\x28\xcf\x84\xd0\xc9\x3a\x27\x42\x7c\x50\x92\xd3\x90\xe6\x42\x4d\x31\x50\x8b\x30\xf1\x14\xc4\x92\xfd\x87\xd8\xfa\xfd\xec\xd0\x6d\xe6\x9c\x25\xee\x3a\xd7\xbf\x2c\x06\xef\xc8\x1c\x17\x3e\x3b\x30\xf7\xa6\x65\x77\x1b\x99\xa3\xea\xd1\x28\x74\xc8\xfc\x65\x3a\xb5\xc4\x5a\x8a\x95\xb3\x43\x3d\x43\x33\x93\x38\x1f\xce\xb5\xb9\x91\x6a\x3a\x20\xc3\x1a\x04\x6d\xdb\x21\x57\x82\xc3\x2f\xf8\xcf\x83\x79\xe1\xfa\x6e\x57\x86\x78\xf3\xdf\xc1\x15\xe1\xb1\xc3\x07\x31\x55\x97\x13\xbb\xc7\xfa\xfd\xf3\x3a\xd1\xdc\x39\x4b\x6e\x31\xcf\x65\x9a\xd7\xb5\x60\x2b\x92\x95\x22\x0b\xa1\x4e\xa8\xc5\x67\x37\x5a\x12\x9d\x37\x84\x7b\x31\x88\xcd\xc7\x40\xa8\x8c\xfe\xb7\xd2\x3a\x5a\x7f\x90\xac\xbc\xdc\xc9\x51\xdf\x9f\x3e\xff\x7b\x4c\xd9\xcb\x07\x79\xe5\x9a\x8a\x88\x9e\x4a\x18\x51\xa2\x43\xd3\x51\x12\x88\x2c\xe3\x66\x4f\x14\x67\x1b\x0b\x87\x07\xe3\x2e\x84\x7a\x71\x8b\xa9\x77\xdb\xcb\xa2\xfd\x0b\x4e\x61\xc2\x20\xb8\xb9\xa6\x80\x4f\x05\x11\x41\x00\xac\x41\x40\x2a\x14\x15\xaf\x4d\xde\x1a\x01\x1c\x1d\x52\x9e\x91\x06\x53\x47\x19\x24\x37\xda\x4f\xf3\x58\xde\x72\x72\x80\x54\x1b\x83\xb6\xd2\x2a\xa3\xb4\xd1\xc8\xa3\x0e\xf4\xed\xba\x30\x39\x6b\xa4\x05\xa5\xa8\x00\x8e\x0f\xe1\x1e\x6c\x8b\x8e\xeb\xf7\x68\x10\xab\xe7\xdb\x1c\xf3\x2f\x0e\x83\x21\xdd\xfc\x5f\x2e\x0b\x6c\xa8\x85\x83\xa3\xc3\x9a\x13\x0b\xb9\xa8\x2a\x54\x96\x92\xb0\x59\x80\x93\x25\x82\x00\x6f\xd1\xc4\xb4\x64\x43\xbe\x0b\xc4\xf5\x41\x2c\xc9\x3a\x38\x3e\x6c\x25\x72\x16\x18\xbb\xaa\xa5\xa6\x21\x6b\x5a\x49\x2b\x9d\x0f\x2d\x3c\xcc\x73\x54\x2d\xbb\x80\x4c\xa3\x55\xfb\xfb\xae\xce\x80\x98\x4c\x13\x42\x87\x46\xea\x4c\xa6\x30\x16\xe9\x8d\xaf\xb8\x7a\x69\xf0\x90\x35\x1b\x99\xd5\x7d\x0c\xde\x4a\xcb\x42\x89\x7b\x27\xb2\xc0\x04\x9e\x35\xf6\x55\x2c\x62\x75\xa3\x99\x4b\xa3\x75\xc9\x94\xa5\x24\xb9\x82\xd3\xb9\x5a\x01\x1a\xbc\x9d\xf8\x33\x5e\x29\x56\x5c\x21\x94\xbd\x93\x9d\xe1\x8d\x76\x38\x82\x15\xa9\x47\x61\xd7\x15\x3e\x0b\x84\x0b\x18\xc2\xb0\xc6\x16\x6c\xa8\x87\x4e\xcf\xe1\xe4\xfd\xbb\x77\x2f\xde\x5c\xbc\xfa\x29\x5a\x1d\xb5\x48\x6f\xb9\x20\x6f\xb5\xe3\xad\xf9\x07\x1c\x9c\x9e\x1c\x92\x68\x32\xad\x30\x94\x3d\x41\x1e\x91\x9a\x7e\xbb\xde\x98\xcb\xa2\x20\xfb\x4d\x0b\x14\x86\x20\xbf\x10\x69\x7e\xd7\xc6\x73\x41\xba\xf6\x4a\x7e\xf0\x08\x14\x78\xac\xae\x2b\x58\xd6\x23\xb1\xc2\x47\xc6\x14\x8c\x06\x4b\x95\x48\x17\x10\x70\x09\x25\x40\xe1\x9c\x8e\xdf\x8f\x26\x9b\x9b\x84\x2a\xda\x6c\x77\x58\xdc\x12\x4e\xa9\x7a\xf6\x9d\x60\xef\x78\x7b\xa3\xad\x73\x3e\x01\xa9\xa8\x48\xa1\xa1\x21\x6b\x1a\x31\x8e\xb7\xba\x28\xb4\x7f\x58\xcf\xb1\x5b\x74\x27\x19\x73\xab\x4e\xd0\x82\x21\xe4\xba\xc8\x6c\xad\x83\xd3\xe7\x71\x06\xd1\x07\xa9\xd2\xc2\xb3\xe9\xbc\x7f\x7f\xfa\xdc\x26\x00\xff\x83\xa9\xf0\x96\xaa\x55\xb2\x80\x7d\x07\x6f\xdf\xbc\xfa\x89\x1c\x37\xec\x88\xea\x27\xf0\x0a\x44\x21\xc3\x24\x24\x10\xcc\xa7\x43\x25\xcb\x98\x1b\x19\x48\xe5\xa8\x83\x27\x7b\xcd\xb1\xa8\x28\x14\xdd\x20\x58\x6f\x22\x75\x04\x98\xdf\x72\xd2\x80\x4c\x73\x85\x3b\x45\x47\x76\x39\x29\xb8\xaf\xff\x0b\x73\xc8\xba\x76\xbb\x43\xd7\xdd\x0d\x77\x50\x71\xbb\xe5\xd6\xe3\x18\x9d\xee\xf5\xdc\x3b\xb6\xdc\x29\x79\x70\x6b\x20\xd9\x7e\xa4\xc3\xb2\xd3\x06\x57\xa8\xeb\x9d\xd4\x20\xea\x46\x87\x48\x74\x42\x16\x96\x23\x13\xb9\xb3\xa0\x2e\xc7\x35\xcd\x53\x08\x49\x6d\xf3\x94\x3c\x98\x83\x7a\x7c\x9a\xc0\x60\x30\x88\x0d\x8e\x33\x9e\x7a\xd4\xa8\xcd\x2c\x46\xe2\x18\xfa\xc9\x3e\x44\xb0\x09\x63\xc4\x02\x44\x18\xdd\x84\x28\x51\x09\x97\x43\x12\xc4\x9b\x2c\x19\x4d\x60\xb5\xc9\x64\x9b\x79\xa9\x75\x14\x6f\x40\xf8\x3b\x33\x3a\x1c\xc2\xbb\x66\x7c\xd4\x12\x78\x8c\xff\x9c\x15\x27\x5a\xef\xdb\x55\x9e\x92\xfa\xf0\x0f\x4a\xcf\x55\x17\x09\x8c\x53\x18\x1c\xc1\x55\xef\xd9\x4c\xc8\x42\x8c\x0b\xbc\xea\xf5\xe1\xaa\x77\x66\xf4\xd4\xa0\xb5\x52\x4d\x69\x81\xcc\xf3\xaa\xf7\x1c\xa7\x46\x64\x98\x5d\xf5\x6a\xd0\x5f\x56\xc2\xa5\xf9\x6b\x34\x53\xfc\x01\x17\x4f\x18\xe0\xca\xab\x73\x67\x84\xc3\xe9\xe2\x49\x49\x7b\x9a\x77\x85\xb4\xee\x62\x51\xe1\x13\x4e\xcb\xad\xc5\xd7\xa2\x5a\x01\xd4\xa8\xd5\xc2\xe5\x75\x89\x4e\xcc\x8e\x92\xa5\xaa\x7f\xf9\xd5\x6a\x35\xba\xea\x2d\x79\xea\xeb\x92\x0c\xa6\x72\x8b\xab\x1e\xac\x50\x30\xba\xea\x31\x0d\xf5\x7a\x4d\xf4\xe8\xaa\x47\xd8\x68\xd9\x68\xa7\xc7\x7e\x32\xba\xea\x8d\x17\x0e\x6d\xff\xa8\x6f\xb0\xea\x93\x5b\x3e\x59\x62\xb8\xea\xfd\x02\x57\xaa\x26\x5a\xbb\x1c\x4d\xd0\xb4\x85\x3f\x7a\x5d\x11\x78\x63\x7c\x06\x28\x84\x75\x17\x46\x28\x2b\xeb\x91\x7a\xf7\xbe\x3b\x06\x7f\xff\x58\x3d\x74\xa6\x37\xa1\xe4\x88\x33\x82\x28\x2c\xd7\xec\x26\xeb\x35\xba\x64\xa7\x08\x56\xc1\x83\x0c\xc5\xcc\xd4\x2d\x7d\xc8\x51\x63\x0c\xf5\x05\x81\xf2\x2a\x43\x53\x2c\x38\x9f\x2f\xbd\x8d\x8b\x8c\x2c\x01\x38\x9d\x84\x78\x16\x5b\xf4\x1b\xb2\x3a\x0a\x98\xa8\xc0\xdb\xba\xb2\x60\xba\x1a\x88\xe4\x6d\xc1\x4b\x22\x18\x0e\xad\x29\x55\xd8\x64\x8a\xdd\x61\x0f\xc8\xf5\x4a\xe1\x46\x40\x45\xe1\x80\x20\xae\xd9\xb7\x25\x9f\x00\x94\x68\xad\x98\xee\x26\xf0\xb8\x37\xe4\x96\xdc\x97\x42\x81\x41\x91\x11\x9d\xcb\x77\x2a\x93\xa9\xe0\xca\xa9\x0e\x3e\x62\xac\x7d\x08\x07\x4b\xf9\x47\x11\x97\x62\x41\xf2\xa5\x72\x81\x0c\xb6\xce\x23\x6b\x88\x29\xc5\xed\x2b\x54\x53\x97\x8f\xe0\xab\xe3\xff\xfc\xe6\xdb\x87\xf2\x5c\x07\xeb\xef\x51\xa1\x11\xdd\xa5\x7c\x07\xfb\xf7\x8f\xb5\x66\xda\xcc\x5f\x52\x8f\x77\x93\xe9\x72\x4f\x18\x4c\xad\xd8\xe1\x5c\x58\x2a\xba\x61\x2c\x2c\x66\xe0\x2b\x92\x07\x85\xc2\xba\x1c\xe3\x2e\xb4\x13\x98\xb4\xad\x42\xf2\xe8\xb8\x0f\xe3\x28\xda\xfb\xb1\xed\xf2\xf6\x3a\xe9\x20\x59\x5a\xf8\xae\x7f\x87\x1e\x2a\x63\x3d\xa7\x05\x6e\x32\xb8\xe8\xa3\xaa\x32\xce\x91\xd6\xe4\x8a\x65\xf9\xb8\xcd\x4a\xa5\x72\xdf\x7c\xbd\x4e\xa9\x52\xc9\xd2\x97\x23\x78\xbc\x51\x9d\x94\x74\xa6\x68\x3a\xf7\x18\x14\x76\x47\x1d\x86\xad\xcb\x04\x29\x28\x38\x4d\x8d\x28\x4b\xe1\x64\x0a\x32\xa3\x16\x7d\x22\xb9\x99\x6a\x0c\x39\x74\xc9\x7c\xb0\x6e\x3e\x1a\xd9\xed\xdb\x18\x6d\x5a\xa6\x7d\x66\x74\xe6\x53\x2a\xc2\xf5\x64\x39\xf2\x6e\x85\xa1\x45\x85\xc1\xf6\x43\x09\x01\x78\x4b\xa2\x6e\xae\x86\xc2\xed\x11\x0a\xea\x12\x6c\x44\x59\x97\xdc\x21\x11\xcd\x73\xe4\xa8\xcb\x17\x5d\xf1\x8c\x61\xaa\xac\xcc\xb8\x76\x16\x30\xf5\xc2\x08\xe5\x10\x33\xbe\x6b\x83\x8b\x7a\x6f\x2b\xb0\x89\xe5\x55\x49\x53\xc3\x5d\x2c\xcb\x7b\x22\x31\x5e\xaf\xb0\x7f\xee\xe0\x98\x47\x8f\x8f\x37\x68\xba\xd9\xb5\x66\x4b\x25\x9c\x43\xa3\x46\xf0\xf3\xe5\xb3\xc1\xff\x8b\xc1\x6f\xd7\x07\xf1\x9f\xc7\x83\xef\xfe\xd1\x1f\x5d\x3f\x6a\xfd\xbc\x3e\x7c\xfa\x1f\x0f\x0d\x01\x9b\xea\xfa\x3b\x26\x13\xd3\xc3\x72\xe0\x1c\xb4\xd8\xe7\xdc\xa1\x27\x70\x61\x3c\xf6\xe1\xa5\x28\x2c\xf6\xe1\xbd\xe2\xa0\xbf\x4e\x50\xa8\x7c\xb9\x0e\xe9\x00\x7a\x04\xaa\x2b\x81\xc6\xd7\x8c\x63\xfd\xfb\x88\xfb\xa1\x22\xe1\x0d\xbb\x08\x84\x2b\x13\x3d\x69\xc7\x8f\xd6\x95\x1b\x8f\xdb\xc9\x51\x74\x12\x2b\xbb\x24\xd5\xe5\xb0\x75\x25\x47\x25\xe5\x6b\xa1\x16\xb0\x0c\x56\xa1\x0e\xbb\x6b\xc9\x36\x74\x0d\xa9\xd1\xd6\x36\x83\x3c\xcb\x37\x07\xd0\x14\x6b\x21\x04\x8e\x63\xd7\x22\xcc\x58\x3a\x23\xcc\xa2\x55\x50\xd7\x43\x16\x6f\x71\xe2\x0b\x38\xb0\x88\x90\x28\x9d\xe1\xfd\x98\x79\x18\x22\xa3\x18\xcb\x42\x3a\x1e\xdf\x67\xc8\x2d\x88\x8c\xa5\x6f\x59\x69\xe3\x84\x72\xc1\x9d\x0c\x4e\xf1\x96\xda\xd4\x92\xca\x29\xe4\xb6\xe8\x20\x53\xf6\xe8\xe8\xf8\xab\x73\x3f\xce\x74\x29\xa4\x7a\x59\xba\xe1\xe1\xd3\x83\x0f\x5e\x14\x7c\xd5\x40\xbd\xe8\xcb\xd2\x1d\xee\x90\xe4\x8e\xbe\xd9\xea\x27\x07\x97\xc1\x1b\xae\x0f\x2e\x07\xf1\xbf\x47\xf5\xd2\xe1\xd3\x83\xab\x64\xe3\xfb\xc3\x47\x44\x5a\xcb\xc7\xae\x2f\x07\x4b\x07\x4b\xae\x1f\x1d\x3e\x6d\xbd\x3b\x7c\xa0\xbb\x19\xfc\xe0\xa5\xc1\xce\xd6\x75\xd0\x51\xc6\x75\x6e\x8b\x05\x46\xe7\xbb\x10\x9c\x3b\x5f\x05\x15\x77\xbe\x22\xaa\xd7\xce\x06\x3b\xc7\x7f\xf5\x4b\xee\x71\x3a\x46\x83\xe7\x6b\xa2\xca\xae\x63\xc9\xd5\x7e\xee\xac\x81\x08\xed\xcf\x16\xf8\x02\xb0\x89\x49\x71\xa2\x78\xa5\xc8\x20\x6d\x18\x37\x86\xdb\x40\xce\xd4\x21\x97\x2c\x93\x93\xad\x0b\xf7\x8e\xe7\x23\xbc\xc1\x19\x9a\x66\x4e\x04\x1f\xd7\xed\x5c\x9e\xd8\xf2\xcc\xc2\x96\x17\xc6\x68\xc3\x07\xfe\x6b\xc0\xcf\x7f\xf3\xf2\x19\x86\xa9\xd5\x0a\xa8\x9f\xb7\xe1\x5a\x87\x74\xb6\x6e\xc3\x97\x01\xe7\xa0\xfe\x3b\xf8\xf2\xcf\xef\xbc\xb7\x3e\x5b\x4b\xd5\x47\x78\x29\x9c\x28\x00\x59\x08\xab\x6c\x9c\x68\x0a\x90\x8e\x1b\xbd\x8f\x0f\x6d\x9d\xde\x87\x2f\x1f\x96\x9f\x22\xfd\xc9\x9e\x41\xf9\x82\x23\xeb\x08\x9c\xf1\x9f\xab\xb1\xd8\x7a\x7e\xfd\xad\xdc\x0e\x87\xab\x5c\xd8\x75\x22\xdb\x30\x4f\x09\xcf\x8a\x17\x9e\x11\xa4\x5d\xbc\x90\xf6\xad\x05\xb9\x4d\x8d\xe1\xd9\x22\xb2\x1d\x79\x0f\xcf\x26\xf1\x7d\x12\xa0\xcd\x95\x52\x78\x3e\xc7\x24\xf4\x93\xc9\xc4\x6a\x0b\x95\x5b\xd4\xde\xc1\xca\xb9\xc3\x6a\x07\xdd\x13\xee\x2d\x60\x77\x33\x80\xf0\xec\x64\x06\xe1\xd9\x51\x38\xe1\xd9\x6e\x12\x0f\x00\xba\x8b\x79\x84\xe7\x73\x1a\xc9\x27\x13\xbe\x31\xc9\xdf\xdf\xd8\x95\xf0\x3f\x09\xd8\x36\x30\x9f\xd0\x8f\xfc\x65\xa2\xdb\x2a\xae\xb5\xb7\x0d\xff\xb6\xf7\x0d\x5b\x85\xb6\xd1\x14\x56\xdb\xca\x42\xa6\x18\xbf\x8d\xa2\xfe\x5a\xf1\x55\x2c\x5f\x19\x53\x43\x33\x27\x76\xf8\xdb\x46\x12\x45\xd8\x5c\x37\x39\xda\x64\x68\x78\xb6\x80\x1f\x7c\xf8\x7a\x41\xc1\x42\x94\x05\x8f\x87\x96\x33\x00\x2b\xa7\x4a\x4e\x64\x2a\x94\x83\x39\x5f\xd5\x32\x78\xe9\xf6\x79\x3e\xf9\x27\x2f\x4c\xee\x2d\x86\x51\x53\xab\x6a\xb0\x4e\x1b\x0a\x64\xad\x15\x3f\x6e\xfa\xb8\xda\xb4\xa2\xe9\xc3\xef\x7f\xec\x2d\xbd\x20\x4c\x40\x43\xb3\xb4\xf2\x69\x74\x2f\x54\x49\xf5\x97\xcf\xfc\xb3\x75\x77\x02\x97\xd7\x7b\x01\x31\x66\x3f\xd6\x9f\x37\xd3\xe2\x3f\x03\x00\x00\xff\xff\x10\xff\x66\xd6\x46\x2e\x00\x00")

func configCrdsKudoDev_instancesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_instancesYaml,
		"config/crds/kudo.dev_instances.yaml",
	)
}

func configCrdsKudoDev_instancesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_instancesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\x1b\x37\x10\xbe\xef\xaf\x18\xa4\x87\x5c\xaa\x55\x8d\xf6\x50\xec\x2d\x70\x73\x30\x5a\xa7\x41\x1c\xf8\x52\xf4\x30\x4b\x8e\xa4\xa9\x77\xc9\xed\xcc\x50\xa8\xf3\xeb\x0b\x92\x2b\xeb\x61\x29\x76\x03\x74\x6f\xfc\x38\xcf\x6f\x1e\xcb\x66\xb1\x58\x34\x38\xf1\x3d\x89\x72\x0c\x1d\xe0\xc4\xf4\x8f\x51\xc8\x27\x6d\x1f\x7e\xd6\x96\xe3\x72\x7b\xd5\x3c\x70\xf0\x1d\x5c\x27\xb5\x38\x7e\x22\x8d\x49\x1c\xfd\x42\x2b\x0e\x6c\x1c\x43\x33\x92\xa1\x47\xc3\xae\x01\xc0\x10\xa2\x61\x86\x35\x1f\x01\x5c\x0c\x26\x71\x18\x48\x16\x6b\x0a\xed\x43\xea\xa9\x4f\x3c\x78\x92\x62\x7c\xe7\x7a\xfb\x43\xfb\x53\x7b\xd5\x00\x38\xa1\xa2\xfe\x99\x47\x52\xc3\x71\xea\x20\xa4\x61\x68\x00\x02\x8e\xd4\x41\x9c\x48\xd0\xa2\x68\xfb\x90\x7c\x6c\x3d\x6d\x1b\x9d\xc8\x65\x67\x6b\x89\x69\xea\xe0\x09\xaf\x2a\x73\x1c\x35\x87\xdf\x67\xed\x02\x0d\xac\xf6\xeb\x11\xfc\x1b\xab\x95\xab\x69\x48\x82\xc3\x81\xb7\x82\x2a\x87\x75\x1a\x50\xf6\x78\x03\xa0\x2e\x4e\xd4\xc1\x87\xec\x6a\x42\x47\xbe\x01\x98\xd3\x2a\xae\x17\x73\xe0\xdb\xab\x9e\x0c\xaf\xaa\x21\xb7\xa1\x11\x6b\x60\x90\x8d\x85\x77\x1f\x6f\xee\x7f\xbc\x3b\x82\x01\x3c\xa9\x13\x9e\xac\x30\xb4\x8b\x11\x58\xc1\x36\x04\x55\x18\x56\x51\xca\x71\x17\x11\xbc\xfb\x78\xf3\x64\x60\x92\x0c\x1b\xef\x48\xa8\xdf\x41\xc9\x0f\xd0\x13\x77\x6f\x73\x44\x55\x0a\x7c\xae\x35\x55\xb7\x73\x6a\xe4\xe7\x24\x20\xae\xc0\x36\xac\x20\x34\x09\x29\x85\x5a\xfd\x0c\x63\x80\xd8\xff\x45\xce\x5a\xb8\x23\xc9\x8a\xa0\x9b\x98\x06\x9f\x9b\x62\x4b\x62\x20\xe4\xe2\x3a\xf0\x97\x27\x6b\x0a\x16\x8b\x9b\x01\x8d\xd4\x80\x83\x91\x04\x1c\x60\x8b\x43\xa2\xef\x01\x83\x87\x11\x1f\x41\x28\xdb\x85\x14\x0e\x2c\x14\x11\x6d\xe1\x36\x0a\x01\x87\x55\xec\x60\x63\x36\x69\xb7\x5c\xae\xd9\x76\xed\xec\xe2\x38\xa6\xc0\xf6\xb8\x2c\x9d\xc9\x7d\xca\xe5\x5d\x7a\xda\xd2\xb0\x54\x5e\x2f\x50\xdc\x86\x8d\x9c\x25\xa1\x25\x4e\xbc\x28\xc1\x86\xd2\xd2\xed\xe8\xbf\x93\x79\x00\xf4\xed\x11\x79\xf6\x98\xbb\x40\x4d\x38\xac\x0f\x2e\x4a\xdb\x7d\x85\xe5\xdc\x7f\xb9\xa2\x38\xab\xd6\x2c\xf6\x64\x66\x28\xf3\xf1\xe9\xfd\xdd\x67\xd8\xb9\xae\x84\x57\x6e\xf7\xa2\xba\xa7\x39\x53\xc4\x61\x45\x52\x25\x57\x12\xc7\x62\x85\x82\x9f\x22\x07\x2b\x07\x37\x30\x05\x03\x4d\xfd\xc8\x96\xeb\xf7\x77\x22\xb5\x5c\x81\x16\xae\xcb\x1c\x43\x4f\x90\x26\x8f\x46\xbe\x85\x9b\x00\xd7\x38\xd2\x70\x8d\x4a\xff\x3b\xc9\x99\x4d\x5d\x64\xf2\x5e\x47\xf3\xe1\x0a\x3a\x15\xae\x3c\x1d\x5c\xec\xf6\xc5\x85\x9a\xec\x06\xed\x6e\x22\x77\xd4\xfa\x9e\x94\x25\xb7\xaa\xa1\x51\x6e\xf0\xa3\x6d\xf2\xf5\x99\x3b\xf5\x72\x72\x75\x31\xb1\xd2\x43\xa9\x27\x09\x64\xa4\x67\xc7\xf6\x45\x6d\x1f\xbf\x45\x6f\x44\x0e\x86\x1c\x48\x9e\x25\x02\xc0\x46\xe3\x19\xf8\x84\xc9\xdb\x27\x13\x33\xde\x93\xe6\xad\xf0\xb4\xca\xf6\x3e\xda\x33\xb6\x2e\x33\x59\x3f\x1a\x91\x87\xf3\x57\x27\x81\xbc\xcf\x92\x65\xcc\x02\xc4\x82\xe1\x50\xd5\x01\xbd\x17\xd2\xb2\x77\x72\xaf\xa2\xab\xc3\x91\x77\xb6\x7f\x21\xbe\x17\x08\xac\x5f\x59\xfe\xaf\x89\x31\xff\x40\xea\x26\x48\x4a\x52\xf4\x20\x0a\x44\x59\x63\xe0\x2f\x75\xa9\x66\xf0\x1b\x23\xb9\x30\x0b\x87\x97\x28\x82\x8f\xcd\xf3\xf0\xcb\x6f\xed\x16\x03\xaf\x48\xed\x3f\xf5\x50\x92\x33\xf5\xb9\x28\x7f\x69\x5c\x0d\x2d\xe9\x6b\x06\xb6\x08\x1e\x8d\x6c\xec\x35\xaf\xc4\x97\x66\xf6\xac\xe7\x67\x60\x35\xd5\x81\x49\xa2\x0a\x58\x14\x5c\xd3\x8c\xec\xe3\x44\xe7\x68\x32\xf2\x1f\x4e\x9f\x1f\x6f\xde\x1c\xbd\x2e\xca\xd1\xc5\xe0\xb9\x3e\x98\xe0\x8f\x3f\x9b\x6a\x95\xfc\xfd\xee\x09\x91\xc1\x7f\x03\x00\x00\xff\xff\x81\x77\x3f\xc2\xaa\x09\x00\x00")

func configCrdsKudoDev_operatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorsYaml,
		"config/crds/kudo.dev_operators.yaml",
	)
}

func configCrdsKudoDev_operatorsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6b\x8f\x1b\x37\x92\xdf\xe7\x57\x14\x26\x1f\x9c\x04\x52\xeb\x9c\xec\xe2\x0e\xf3\x2d\x6b\xc7\x87\xb9\x78\x6d\xc3\x8f\x3d\x1c\x72\x01\x86\x6a\x56\xab\xb9\xc3\x26\xfb\xf8\x90\xac\x0b\xf2\xdf\x0f\x55\x64\x3f\xa4\xd1\xa3\x2d\x27\xb9\x2c\x76\xfa\x4b\xa2\x6e\xb2\x58\xef\x17\x6b\x7c\x35\x9f\xcf\xaf\x44\xab\xfe\x86\xce\x2b\x6b\x6e\x40\xb4\x0a\x3f\x06\x34\xf4\xcb\x17\xf7\xff\xe6\x0b\x65\x17\xeb\xa7\x57\xf7\xca\xc8\x1b\x78\x16\x7d\xb0\xcd\x5b\xf4\x36\xba\x12\x9f\x63\xa5\x8c\x0a\xca\x9a\xab\x06\x83\x90\x22\x88\x9b\x2b\x00\x61\x8c\x0d\x82\x5e\x7b\xfa\x09\x50\x5a\x13\x9c\xd5\x1a\xdd\x7c\x85\xa6\xb8\x8f\x4b\x5c\x46\xa5\x25\x3a\x06\xde\x1d\xbd\xfe\x97\xe2\x4f\xc5\xd3\x2b\x80\xd2\x21\x6f\x7f\xaf\x1a\xf4\x41\x34\xed\x0d\x98\xa8\xf5\x15\x80\x11\x0d\xde\x80\x6d\xd1\x89\x60\x5d\xde\xe9\x8b\xfb\x28\x6d\x21\x71\x7d\xe5\x5b\x2c\xe9\xcc\x95\xb3\xb1\xbd\x81\xfe\x7d\xda\x99\xd1\x49\xa4\xbc\xce\x40\x32\xe5\xfc\x45\x2b\x1f\x7e\x38\xf4\xf5\xa5\xf2\x81\x57\xb4\x3a\x3a\xa1\x1f\xa2\xc0\x1f\xbd\x32\xab\xa8\x85\x7b\xf0\xf9\x0a\xc0\x97\xb6\xc5\x1b\x78\x45\x68\xb4\xa2\x44\x79\x05\xd0\x6d\x26\xb4\xe6\x99\xb6\xf5\xd3\x25\x06\xf1\x34\xc1\x2b\x6b\x6c\x44\x42\x1a\x08\xa6\xf9\xee\xcd\xed\xdf\xbe\x7d\xb7\xf3\x1a\x40\xa2\x2f\x9d\x6a\x03\x33\x71\x0f\x71\x50\x1e\x42\x8d\x90\xf6\x40\x65\x1d\xff\xdc\x47\x1f\xbe\x7b\x73\x5b\xf4\x00\x5b\x47\xdf\x83\xea\x18\x96\x9e\x91\x96\x8c\xde\xee\x1d\xff\x84\x30\xcc\x47\x4b\x52\x0f\x4c\xe7\xe7\x83\x50\x66\xa2\xc0\x56\x10\x6a\xe5\xc1\x61\xeb\xd0\xa3\x49\x0a\x43\xaf\x85\x01\xbb\xfc\x3b\x96\xa1\x80\x77\xc8\x18\x82\xaf\x6d\xd4\x92\xf4\x68\x8d\x2e\x80\xc3\xd2\xae\x8c\xfa\xdf\x1e\x9a\x87\x60\xf9\x18\x2d\x02\xfa\x00\xca\x04\x74\x46\x68\x58\x0b\x1d\x71\x06\xc2\x48\x68\xc4\x16\x1c\x12\x5c\x88\x66\x04\x81\x97\xf8\x02\xfe\x6a\x1d\x82\x32\x95\xbd\x81\x3a\x84\xd6\xdf\x2c\x16\x2b\x15\x3a\x0b\x28\x6d\xd3\x44\xa3\xc2\x76\xc1\xca\xac\x96\x31\x58\xe7\x17\x12\xd7\xa8\x17\x5e\xad\xe6\xc2\x95\xb5\x0a\x58\x86\xe8\x70\x21\x5a\x35\x67\x64\x0d\x5b\x41\xd1\xc8\x2f\x5c\xb6\x19\xff\x64\x87\x79\x61\x4b\x5a\xe1\x83\x53\x66\x35\xfa\xc0\x2a\x7a\x82\xcb\xa4\xa4\x24\x5a\x91\xb7\x26\x2a\x06\x66\xd2\x2b\xe2\xc7\xdb\xef\xdf\xbd\x87\xee\xe8\xc4\xf0\xc4\xdb\x61\xa9\x1f\xd8\x4c\x2c\x52\xa6\x42\x97\x56\x56\xce\x36\x0c\x05\x8d\x6c\xad\x32\x81\x7f\x94\x5a\xa1\x09\xe0\xe3\xb2\x51\x81\xe4\xf7\x3f\x11\x7d\x20\x09\x14\xf0\x8c\x4d\x1f\x96\x08\xb1\x95\x22\xa0\x2c\xe0\xd6\xc0\x33\xd1\xa0\x7e\x26\x3c\xfe\xe6\x4c\x26\x6e\xfa\x39\x31\x6f\x1a\x9b\xc7\x5e\x6b\x7f\x71\xe2\xd3\xe8\x43\xe7\x5b\x8e\xc8\x64\xcf\xf0\xde\xb5\x58\xee\x58\x80\x44\xaf\x1c\x69\x6c\x10\x01\x49\xcf\xf7\x36\x14\x3b\xa0\x0f\x9b\x60\x32\xc3\xf6\xa0\x19\x9e\x20\x13\x92\x0f\x36\x58\x12\xaa\xef\xf8\xf3\xc3\xcd\x3b\xd4\x3c\xdb\x5b\xde\x93\x22\x20\x60\xd3\x92\x9d\xc9\x4e\xf7\x42\x2d\x02\x94\xc2\xb0\xdc\x3d\x4a\x32\xc6\x7c\x1c\xfd\xaf\x30\xa0\x8c\x0f\xc2\x94\x98\xac\x1e\x7b\xd2\x8b\x4f\xa1\xa0\xf3\x59\x67\x30\x7f\xf2\x9a\x05\xf7\x16\x2b\x74\x48\x67\x92\x2e\x09\x65\x3c\xa0\xb1\x71\x55\xb3\xfa\xb9\x26\xb9\x9b\x60\x41\x63\x80\xad\x8d\x84\x63\x4b\x18\x5b\x07\x8d\x95\xaa\xda\x32\xa6\x8e\xc0\x90\xd8\x3a\x97\x34\x9f\xcf\xe1\x15\x6e\x88\x50\xdf\x3b\x31\xc2\x1a\x84\x43\x90\xca\x97\x36\x3a\xb1\x42\x09\x4b\x2c\x45\xf4\x4c\xb3\x54\x55\xa5\xca\xa8\xc3\x36\xe3\xba\x24\xbe\x91\xf9\x44\x2f\x56\x08\x9b\x1a\x0d\x60\xb3\x44\x29\x51\x82\x32\xe4\x8e\x7d\x01\xf0\xb4\x80\xdb\x95\xb1\x74\x7e\xa5\x50\x4b\x7a\x77\x4b\xee\xad\xd4\x51\x22\x19\xac\xd9\xe6\x2f\xb0\xa9\x55\x59\x33\x12\x64\x82\x2b\x34\xe8\x84\xd6\x5b\xa8\x2d\x03\x28\x00\x5e\x58\xd7\x4b\x62\x06\x5d\x10\xef\xbc\x35\xf9\xc8\x17\x04\xea\x8d\x08\x09\xce\xd2\x86\x9a\x1c\xf7\x16\x9c\x70\xa8\xb7\xe4\x64\x14\xa3\x27\xca\x10\x85\x4e\xc8\x17\x00\xdf\x90\x99\xa7\x8f\x89\x9e\x1a\x75\x9b\x51\xf5\xa0\x9a\xd6\x7a\xaf\x96\x1a\x59\x1b\xa4\x64\x4b\x52\x95\x2a\x79\x1d\xc7\x24\x65\xa4\x5a\x2b\x39\x06\x7a\x6b\xa0\xb1\x3e\x0c\x6c\xe1\x0f\x7e\x46\x62\x71\x89\xdb\xad\x70\x81\xd8\x2a\x1c\xab\x81\x43\xd2\x1b\x56\x5a\x0f\x5a\xdd\xe3\x0c\xae\x9b\xe8\x43\x12\x22\x58\xa3\xb7\x1c\x27\xc8\x49\xc0\x77\x4c\xf0\x5f\xae\x49\xde\xd7\x1f\x6e\x9f\x33\xd7\x32\xaf\xd2\x4b\x8a\xc7\xc0\xfb\x97\xd8\xc3\x46\x79\x5d\xf0\x61\xef\x6b\xeb\x91\xb4\x3e\x3b\xbc\x0d\x6a\xdd\x09\x17\xe5\xae\x44\x0b\x80\x6f\x89\x45\xa5\x35\x5e\xf9\x40\xee\x93\xb9\xc5\x3a\x58\x00\xfc\x25\x6b\x0a\x29\x5c\xa2\x32\x2b\x53\xc5\x3a\x1c\x66\x29\x84\xf6\x5b\xc0\x45\xbd\xbf\x06\x96\xdb\xb4\x77\x96\x35\xa1\x11\xf7\xe8\x41\x05\xa8\x85\x93\xcc\xe4\xe8\xc9\xc9\x07\x0b\xad\x43\xa9\xca\x00\x1b\x32\xdc\x8d\xd2\x1a\x6a\xd1\xb6\x48\xa8\xfc\xa9\x80\xf7\x35\x76\x3a\xd5\x6b\x81\x6a\x5a\x87\xa5\xf2\xc8\x5c\xb3\x6b\x74\x7a\x0b\xf9\x55\x01\xd0\x85\x23\xe2\x85\xe8\xde\x43\x23\xda\x96\xfd\x83\x05\x01\x1f\xde\xbe\x24\xd0\xca\xb3\xa7\x68\x9d\x95\xb1\x44\x10\xcd\x52\xad\xa2\x0a\xdb\x64\xc7\x91\xfd\x09\x47\xef\xd6\x61\x4e\x09\xe8\x44\x8a\x32\x8a\xa4\x9e\x22\x5a\x86\x3c\xd2\x92\x52\xf8\xac\x1b\x20\xb1\x45\x23\xd1\x94\x5b\x42\x89\x8c\xbc\xc6\x94\x10\xce\x86\x48\x18\x5b\x8d\xc9\x9d\x1a\x39\x4e\x50\x3a\x0f\x95\x35\xdc\x07\x17\xcb\xa4\xc5\xce\xa1\xc6\xb5\x30\xa1\x00\xf8\x73\x01\xff\xd9\x0b\x1f\x85\x57\x7a\x0b\x65\x2d\xcc\x0a\x41\x85\x1d\x81\x76\xce\x41\xf9\x1d\xfb\x66\xc3\xd5\xb6\x4c\x59\xf2\x2c\x87\xcb\x9c\xc6\x74\x7b\xe8\x61\xe9\x88\xaa\x22\xcf\x64\x62\x83\xce\x46\xdf\x25\x3d\x05\xc0\x73\x6b\x9e\x3c\x09\x2c\x6b\x30\xb8\x61\xbf\x91\x0e\x22\xb7\x1b\x8d\x44\x97\x8d\x0d\x25\x7d\x4c\x80\x43\x8d\x5b\x90\x96\xc5\x95\x73\x73\x52\x4f\x1f\x50\x48\x62\x40\xf4\xc9\xad\x67\x44\x66\x29\x21\x27\xee\x13\xca\x9a\x45\x6f\xd7\x4a\xf2\x29\x32\xfb\xfc\x04\x58\x30\xb3\xc8\x18\xe6\x95\x2d\xf9\x8b\x35\xe4\x5f\x5d\xb2\x42\xf2\xc8\x05\x7b\x22\xfc\x28\x9a\x56\xe3\x8c\xb3\x0f\x55\x62\xef\xb0\x3d\x2b\xab\x90\x8d\xf2\x2c\x11\x87\x2b\xe5\x83\x13\xc9\xbd\x8f\xd2\x86\x3a\x2e\x8b\xd2\x36\x0b\xaa\x27\x9c\xc1\x80\x9e\x72\x82\xc5\x52\xdb\xe5\x82\x84\x25\x3c\xce\x9f\x16\x4f\xff\x75\xd1\xc3\x1a\x83\x5a\xac\x9f\x2e\xd8\x15\x14\x2b\xfb\xc5\xcb\x3f\x7f\xfb\x2d\x14\x4f\x1e\x44\x96\xe3\x61\x18\x4e\x64\xc4\x07\xe3\x12\x71\x7f\x4f\xc9\x32\x47\xc2\xc3\x30\x78\x26\x14\xd2\x53\x75\xbe\x7a\xc2\xd9\x4f\x6e\xab\x1c\xc9\x7a\x7b\x6c\x15\xa6\x78\xdc\xa7\xdb\x1c\x1b\xb2\x06\x08\x03\x94\x56\x39\xcc\xdf\x66\x49\x1b\x72\xc0\x1f\xd2\x71\x0a\xac\x20\x72\x60\xf8\x8f\x77\xaf\x5f\x2d\xfe\xdd\x26\xcc\x40\x94\x25\x7a\x9f\xd2\x9d\x86\x9d\x98\x8f\x14\xa0\x7c\x97\x09\xbd\xa3\x2f\x45\x23\x8c\xaa\xd0\x87\x22\x43\x43\xe7\x7f\xfc\xe6\xa7\x3d\x15\x51\x89\x5f\x7d\xea\xda\x85\x76\xe5\x13\x31\xfd\x5e\xd8\xa8\x50\x33\x4a\xad\x95\x19\xe9\x0d\x23\x1b\xc8\x44\x6c\x46\x36\x22\xc7\x87\x1b\xb8\x26\xeb\x18\x1d\xfd\x33\x39\xfd\x5f\xae\xe1\xcb\x0d\x07\x19\x8e\x01\xd7\xe9\xc0\xbe\xc6\xe0\xb8\x90\x25\x38\x1c\xcc\xaa\x1f\x9c\x5a\xad\x90\xc2\x35\xa7\xcd\x94\x9a\x7e\x45\xb1\x44\x55\x60\xec\x68\x31\x83\x20\x7e\xf6\xb6\xb9\x8f\xc8\x8f\xdf\xfc\x74\x0d\x5f\xee\xd2\x45\x51\x12\x3f\xc2\x37\xe4\x40\x98\xb2\xd6\xca\xaf\xb2\x53\xf5\x5b\x13\xc4\x47\x82\x59\x52\x60\x32\x7d\xb4\xab\xc5\x1a\xc1\xdb\x26\x45\xa8\x79\x4a\xe3\x24\x6c\xc4\x96\x68\xe8\x58\x49\x52\x15\x1c\x4f\xf7\x2a\xb0\xf7\xaf\x9f\xbf\xbe\x49\xa7\x91\xd8\x56\xa6\x73\xf3\x95\xa2\xfa\x2a\x79\x4f\xaa\x15\x58\xe6\x84\x48\x4c\x42\xa2\x1c\x30\x7b\xc4\xe4\x81\xab\x48\x59\xfb\x01\x1b\x9b\xa0\xeb\x0f\xcb\xa1\xc3\x6a\xce\x71\x68\xdf\xb8\xfe\xdf\x8a\x8e\x89\xc4\x71\xdd\x3f\x81\xb8\x57\x23\xbd\x3b\x49\xdc\xe0\x0f\x89\x3e\x69\x4b\x4f\xa4\x95\xd8\x06\xbf\xa0\xd0\xbd\x56\xb8\x59\x6c\xac\xbb\x57\x66\x35\x27\xc5\x9a\x27\x69\xfb\x05\x37\x49\x16\x5f\xf0\x7f\x2e\xa6\x85\xdb\x1b\x53\x09\xe2\xc5\xbf\x07\x55\x74\x8e\x5f\x5c\x44\x94\xdb\xcd\x94\xa7\x90\xf6\xae\xcb\x70\xf7\xf6\x92\x59\xa4\xf4\x2c\x37\x3f\x46\x9e\xac\x11\x32\xb9\x3a\x61\xb6\xbf\xb9\xd2\x12\xeb\xa2\xa3\xb3\xb7\xf3\x9c\x02\xcc\x85\x91\xf3\x3e\x45\x2d\xb7\x17\xf1\x2a\xaa\x49\x86\x4a\x09\xf7\xef\xa2\xca\x51\x5d\x64\x95\x47\x5a\x00\xf4\xb4\xc2\x89\x06\x03\xba\x03\x29\x81\x0a\xd8\x1c\xcc\x14\x76\xa8\x7f\xd3\x41\x80\x52\xb4\x24\xa0\xdc\x22\x13\x4e\x89\xa5\xd2\x94\x0d\x27\x27\xbc\xdf\xcb\x5b\x62\x4a\x8f\xa9\x84\x0b\x8a\x4b\x70\x8a\x75\x43\x7d\x7d\x28\x91\x38\x9d\xc2\x10\x6a\x95\x88\x3a\x1c\xfe\xb8\x87\xf9\xf3\xb4\x36\x75\x9e\xf2\xc6\x1c\x4f\x53\x88\xeb\x99\x43\x4b\xfa\x24\x71\x99\x6a\xe9\x53\x58\x9e\x95\xc8\x3e\x2e\xd3\xd0\xed\x7f\x0c\xac\xa6\x24\xd6\xac\xd0\x8d\x97\x12\xbf\x6b\xbb\x61\x2c\x07\x12\x38\xf7\xce\x3d\x8d\xcb\x71\x56\xbe\xd5\x62\xfb\xea\xa8\x93\xdf\xc7\x79\x58\xbf\xd3\x53\x59\x6e\xe1\xc3\xad\xbf\x18\x0d\x34\xb1\x99\x2a\xe2\xdc\xe7\xd1\xca\xa7\x6c\x40\x6b\xbb\x19\x35\x4a\x6f\xab\xb1\x1e\x78\x0c\x9c\x05\x7c\x6f\x62\xd3\xe5\x06\x46\xe9\xbe\x64\x8d\x43\x0d\xdd\xa5\x2d\x0c\x58\xa4\x2a\xe1\x08\x4a\x47\x0d\x69\x22\xb9\xdd\x12\xe1\x9c\xd8\x1e\x5c\xa1\x9a\x26\x06\xb1\xd4\xd3\xa4\x92\xfd\x39\x15\xd4\xd5\x9e\x96\x64\x21\xa5\x64\x47\x82\xa8\x02\xba\xac\xee\x2a\x28\xa1\x93\xda\x6b\xdd\xf7\xb7\xc7\xfd\xf7\x93\xc8\x2f\xad\xd5\x28\xcc\xc1\x35\xc7\x93\x86\x3d\xcc\xaf\x5f\xe5\x5c\x93\x8e\x1d\x37\xec\x72\x12\xdf\xe9\x57\xce\xd2\xba\xe6\x1e\x54\x4a\x23\x17\x62\xe3\x24\xfc\x2e\xdd\x51\x3c\x7b\xfd\xe1\xd5\xfb\x3b\x5a\x6f\xfa\x5a\xb1\xf3\x5f\x9a\xe5\x2c\x38\xb5\xcd\x49\xf6\x7f\x9b\xd4\x3b\xe5\x50\xda\x6a\x55\x0a\x7f\x03\xf0\xf3\xcf\x50\xb0\x27\xf4\x05\xc3\x83\x5f\x7e\xb9\xbe\x54\xbb\x73\x7b\xe0\x48\xe8\xd9\xe3\xc8\xdb\xbc\xb8\xcf\xbe\x0f\x08\x55\xf9\x1e\x26\x85\xec\x25\xee\x38\x33\xa1\x75\xef\xcc\xfc\x8c\x12\xfc\x4d\x8d\xa1\x46\x37\xf2\x8a\xa4\x16\x3e\x56\x95\x3a\xe7\xef\x4e\x49\x39\xd7\x13\x93\xc8\x7a\x9f\xd6\x82\x92\x14\xe6\x99\x2c\xa6\x49\x0b\x93\x04\xbe\xc2\xe0\x01\x3f\x62\x19\x43\xd7\xa0\x4a\x55\xc4\xa0\xca\xac\xc3\xbe\xd3\x85\xdb\xbe\x6b\x9b\x8b\x81\x91\xd9\xdf\xa5\x8e\xc5\x1d\xe7\x2b\xe9\x10\x2e\x51\xf8\x24\x2e\x6f\xf0\xa3\xf2\x81\xb8\x43\x8c\xd9\x28\x8f\xa0\xc2\x13\x0f\x77\x12\x5b\x6d\xb7\x77\x17\x7b\x32\xf6\x29\x73\x5e\x36\x89\x2d\xdb\x16\x47\x92\x1e\xbc\x12\x41\xe8\x49\xe2\xe2\xee\x2e\x9d\x7a\x29\x6a\x27\x72\x86\x53\xee\x88\x78\x77\xc0\xd5\x09\x29\xf9\x66\x55\xe8\x37\x27\x03\xf8\x6e\x66\x41\x72\x18\x88\x15\xe0\xd1\xa9\xd4\xa7\x7e\x53\x0b\xcf\xf4\x93\x7c\xb0\x57\xeb\xd2\x92\x71\x87\xc3\x01\xee\x5c\xea\xd0\x32\xcc\x49\x62\xc8\xc7\x37\xa2\x25\xb4\x78\x63\x52\x13\xae\xdf\xf9\x6b\xa7\x66\x97\x85\x84\x87\xa7\xed\x30\xa2\x0b\x66\x3e\x60\x9b\xb9\xd0\xb5\x2f\x7e\xe8\x73\xcc\x8c\xc1\xd1\x28\x7b\x9e\x23\xe9\x39\xe5\x9c\xd3\x33\x21\x80\xd1\xc3\xd8\x9e\x86\xb4\x1b\xa6\x98\xba\xcc\x64\xda\x3c\xe2\x71\xc7\x81\xe1\xfa\xe6\x21\xe1\xe0\x03\x5f\x36\x88\xe1\x66\xf1\x38\x2f\xe0\xbc\x50\x8e\xa0\x38\xba\x4c\xea\xaf\x02\x28\x8f\xb0\x55\x0e\x8c\xdc\xc1\x67\x21\xd9\xb2\x8c\x07\xae\x89\x76\x9f\x69\x52\x49\xcf\x79\xd9\xa4\x67\xa2\x84\xf2\x62\xe1\xef\x27\x9c\x3d\x91\x5f\x17\x20\x70\x3e\xdf\xd9\x5f\x79\xc4\x4f\x7d\x3a\x48\xee\xa7\xe2\x6a\xfb\x09\x6a\xfa\xda\x49\x4c\x8d\xc4\xde\x42\xbb\xec\xdb\xc7\x25\xf3\x68\xe8\x71\x69\x61\x16\xc9\x5f\x0c\x19\x0b\x0f\x93\x48\xb0\xf1\xa8\xb7\x18\x53\x70\x86\x87\x13\xf8\x61\xa2\xd6\x9c\x2b\x42\x70\x11\x2f\xcc\x38\xcf\x31\xea\xf7\x66\xd1\xe7\x84\xb3\x07\x51\x67\x70\xed\x14\x82\x06\xaf\x43\x3f\x1f\x22\x70\x86\x9f\x27\x8e\x3e\x62\x69\x3b\xf8\xbc\x1c\x0a\x97\xb4\x1e\xc4\x5a\x28\xdd\xe5\xa6\xcc\xb3\x93\x97\xef\x30\xb9\x94\x7f\x2f\xfc\x7d\xaa\x86\x57\xda\x2e\x85\x9e\x41\x6b\xf5\xb6\xb1\xae\xad\x55\x09\x8a\x62\x6b\xb3\x33\xdb\xa2\x35\xb4\x71\xa9\x55\xa9\xb7\x23\xac\x18\xcb\x0b\x02\xf0\xf1\xbe\xe8\x04\xdd\x3f\xe5\x09\xcf\x6e\x7e\x38\x0c\x71\x82\x43\x3c\x0b\xc1\xb7\x52\x3e\x71\xa0\xbf\x0f\x26\xf6\x11\x28\x9f\xdb\xd8\x5c\x6b\x7a\x88\xa9\xaf\xbf\xb6\x4a\xc2\xc6\x29\x1e\x67\x29\x79\xf4\x0c\xa2\x59\x34\xc2\xf9\x5a\x68\xcd\x3d\x79\xbe\xc2\x64\xa5\xe7\x8e\x77\x2b\x9c\x47\x28\xd1\x71\x68\xcf\xd7\x98\xe9\x46\x90\x80\xe4\xdb\x40\x3e\xf7\x07\x65\x64\xba\xed\x94\x76\x63\xbc\x92\xd8\xdf\xe7\x8b\xb6\x75\x56\x94\x35\x28\xbe\x53\x14\xa3\x5b\xe8\x74\x7b\x4c\xd9\x3d\x5f\x18\x8b\x75\x7f\x59\x9a\xf3\x5f\x04\x4f\xda\xff\x77\x6f\x93\x1d\x78\x8a\xa6\xaa\x43\x72\x89\xa5\x6d\xba\x7b\x4f\x1b\x7d\x3f\x91\xd5\xd5\x0d\x4c\x80\xe3\xfb\xc5\x46\xad\xea\x00\x0e\xd7\xca\xab\xb0\x8f\xd8\xb8\xa9\xde\x99\x3d\x2f\xe9\x4e\x30\xa0\xbc\x8f\x47\x8b\x8f\x29\x31\xf3\xd4\xa8\xc9\x11\x71\x8f\x22\xba\x68\xdb\xfe\xc2\x2b\xa3\x6b\xa9\x1e\xa2\x9a\xd8\x61\x6b\x67\x1d\xcd\xfd\xcd\x0a\xdf\xe4\x3a\x2c\xd1\x9c\x0a\x4c\x93\x9c\xba\xb4\xe6\x64\x98\x3f\x5f\x7b\xd1\x53\x89\x20\xf4\xe7\x83\xe9\x6a\xc5\x53\x4d\x20\x98\x4a\x99\xdd\xf5\x5c\x97\x48\xa6\x03\x71\xb9\x78\xc0\x9a\x63\x61\x70\x32\x21\xad\x28\xef\xc5\xea\x24\x43\x76\x08\x40\xc5\x25\x36\xe1\xd6\xed\x65\x03\x9b\xa5\x5b\xe9\xfe\x5d\x65\xb5\x44\x47\x55\xb9\x30\xf0\xe1\xed\x4b\x9e\x78\xc8\xdf\x82\x70\x4b\xa1\x75\xd1\x8d\x1a\xf4\x9c\x18\x37\x69\x66\x3c\x37\x5a\x06\x9d\x3a\x80\x0e\xbd\xd5\x6b\xcc\x1d\x82\x04\xa7\x9b\x82\x70\xe4\x35\x46\x17\x60\xbd\x0f\xc8\x9b\xe4\x70\x02\xa1\x7a\x2a\x5b\x99\xc8\xb3\x5c\xab\x7f\xb6\x1a\xf5\x90\x5e\xa8\xe3\x3d\xb0\x07\x32\x18\xdf\x7d\xee\xb6\x8a\xbe\xe4\x69\x86\xa1\x91\x74\xd7\x7d\xf6\x77\x59\x22\x5f\xa5\x81\xc2\xee\xfa\x03\xe1\xeb\x56\x38\x34\xe1\xeb\x61\x3c\x2c\x4d\x2d\x05\xae\x0c\x86\xb6\x04\xc3\xef\x06\xcb\x5a\xdb\x46\x3e\x95\x21\x94\xb5\xd2\xf2\xeb\xbe\x53\x51\x50\xa4\x29\xfa\xfe\xfa\xa1\x90\xfa\x69\x4c\x52\xc7\xbb\x0c\x30\x2d\x9d\xdf\xcd\x94\x54\x8b\x79\x30\x30\x0d\x0a\xa5\xac\x4e\x24\x12\x3b\xea\x53\x8f\x89\x17\x77\xf9\x45\x57\x97\xf1\x28\x4c\xbe\x33\x92\x27\xcf\x9d\x5e\x13\xa1\x59\x9f\xd3\x81\xf4\x7c\x42\x45\x52\xfd\xda\x00\xef\xf1\x4c\x81\xf1\xa9\xf0\x4e\x64\x4e\x17\x00\x9c\x58\x53\x4d\xaa\x23\x60\x62\xe9\xd5\xda\x93\x14\x4c\xc2\xbd\x9f\x4a\xfe\x4c\x2d\x9f\xc8\xa8\x5f\x95\xfe\x8d\x30\xe1\x7b\x37\xc1\x13\x9e\x0e\xcf\x67\x44\x77\x61\x57\xaf\xf7\x7e\x97\x77\xf6\x4e\xf2\x74\x37\xc7\xee\x0e\x4b\xa5\x48\xd7\xe9\x71\xc3\x44\x55\xb0\xf0\x5f\xdf\xfd\xf5\xe5\x80\x16\xec\xb9\xeb\xe1\x43\x0e\x9f\xe4\x6a\xe8\xc5\x68\x6e\x4b\xe6\x89\x70\x2a\x56\x8e\x4d\xee\x1e\x64\x54\x6c\x57\x4e\x48\x12\xfc\x0b\x67\x0f\x5c\x43\xed\x10\xf3\x61\x67\x31\x13\x93\xaa\x86\xbd\x92\xcd\x0f\x73\xc7\x09\x3e\xf6\xd3\x73\xbf\x52\x71\xf7\x38\x49\xfc\x38\x49\xfc\x38\x49\xfc\x38\x49\xfc\x38\x49\xfc\x38\x49\xfc\xd9\x93\xc4\xe7\x33\xf2\x73\xd3\xc4\x9f\x3b\x4f\x3c\x21\x4b\x3b\x33\x53\xfc\x38\x55\xfc\x38\x55\xfc\x8f\x34\x55\x3c\x41\xe3\x4f\xd5\x81\xff\x08\xb3\xc5\x9f\xd9\xe7\xff\x03\x4e\x18\x4f\xa4\xe8\xc4\x94\xf1\x1f\x76\xce\x78\xd2\x24\xd3\x84\x59\xe3\x7f\x9e\x69\xe3\x09\x1c\x3b\x3a\x71\xfc\x07\x9c\x39\xfe\xad\x66\x88\xd6\x9f\xfc\x07\xc1\xc7\xfe\xc6\x39\x88\x10\xfd\x27\xfc\x95\x33\xaf\xdf\xf9\x3b\x67\xbb\xf4\xe8\xd6\x93\xff\xd0\xf9\x20\x22\x0f\x5e\x26\x90\xa3\xb6\x91\x0f\x96\x6a\xe3\xfc\x66\x40\x9b\xf2\x83\x36\xa0\x7c\xb5\xff\xcf\x3d\x5c\xa7\xe1\xc2\xee\xdf\x6f\xe0\x9f\xa5\x35\xa9\x15\xe3\x6f\xe0\xc7\x9f\xae\x20\xb7\x59\xbb\x06\x03\xbf\xfc\xbf\x00\x00\x00\xff\xff\x8f\xe3\xc6\x4f\x21\x43\x00\x00")

func configCrdsKudoDev_operatorversionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorversionsYaml,
		"config/crds/kudo.dev_operatorversions.yaml",
	)
}

func configCrdsKudoDev_operatorversionsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorversionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/crds/kudo.dev_instances.yaml":        configCrdsKudoDev_instancesYaml,
	"config/crds/kudo.dev_operators.yaml":        configCrdsKudoDev_operatorsYaml,
	"config/crds/kudo.dev_operatorversions.yaml": configCrdsKudoDev_operatorversionsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"kudo.dev_instances.yaml":        &bintree{configCrdsKudoDev_instancesYaml, map[string]*bintree{}},
			"kudo.dev_operators.yaml":        &bintree{configCrdsKudoDev_operatorsYaml, map[string]*bintree{}},
			"kudo.dev_operatorversions.yaml": &bintree{configCrdsKudoDev_operatorversionsYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
