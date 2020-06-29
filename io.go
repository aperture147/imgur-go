package imgur_go

import "io"

type Image struct {
	FullName    string
	Description string
	Data        io.ReadCloser
}

func (i Image) Name() string {
	return i.FullName
}

func (i Image) Read(p []byte) (n int, err error) {
	return i.Data.Read(p)
}

func (i Image) Close() error {
	return i.Data.Close()
}

type NamedReadCloser interface {
	Name() string
	io.Reader
	io.Closer
}
