package imgur_go

import "io"

type Image struct {
	Info ImageInfo
	Data *io.Reader
}

func (i Image) Name() string {
	return i.Info.Title
}

func (i Image) Read(p []byte) (n int, err error) {
	return (*i.Data).Read(p)
}

type NamedReader interface {
	Name() string
	io.Reader
	io.Closer
}
