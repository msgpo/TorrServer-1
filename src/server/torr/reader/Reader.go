package reader

import (
	"github.com/anacrolix/torrent"
	"io"
)

type Reader struct {
	torrent.Reader
	offset    int64
	readahead int64
	file      *torrent.File
}

func NewReader(file *torrent.File) *Reader {
	r := new(Reader)
	r.file = file
	r.Reader = file.NewReader()
	return r
}

func (r *Reader) Seek(offset int64, whence int) (n int64, err error) {
	switch whence {
	case io.SeekStart:
		r.offset = offset
	case io.SeekCurrent:
		r.offset += offset
	case io.SeekEnd:
		r.offset = r.file.Length() - offset
	}
	n, err = r.Reader.Seek(offset, whence)
	r.offset = n
	return
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.offset += int64(n)
	return
}

func (r *Reader) SetReadahead(length int64) {
	r.Reader.SetReadahead(length)
	r.readahead = length
}

func (r *Reader) Offset() int64 {
	return r.offset
}

func (r *Reader) Readahead() int64 {
	return r.readahead
}
