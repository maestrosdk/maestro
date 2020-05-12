package diagnostics

import (
	"io"

	"compress/gzip"
)

type streamGzipper struct {
	bufSize int
	w       io.Writer
}

func newGzipWriter(w io.Writer, size int) *streamGzipper {
	return &streamGzipper{
		bufSize: size,
		w:       w,
	}
}

func (z *streamGzipper) Write(r io.ReadCloser) error {
	buf := make([]byte, z.bufSize)
	zw := gzip.NewWriter(z.w)
	var err error
	for {
		var n int
		n, err = r.Read(buf)
		if n > 0 {
			_, err = zw.Write(buf[:n])
		}
		if err != nil {
			_ = zw.Close()
			_ = r.Close()
			break
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
