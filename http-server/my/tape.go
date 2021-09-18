package main

import (
	"io"
)

type ReadWriteTruncate interface {
	io.ReadWriteSeeker
	Truncate(int64) error
}

type tape struct {
	file ReadWriteTruncate //*os.File //io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (int, error) {
	t.file.Seek(0, 0)
	t.file.Truncate(0)
	return t.file.Write(p)
}
