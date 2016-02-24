package model

import "os"

// todo: docs
type File struct {
	Path string
	ref  *os.File
}

// todo: docs
func NewFile(path string) (*File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0664)

	if err != nil {
		return nil, err
	}

	result := &File{
		ref:  file,
		Path: path,
	}

	return result, nil
}

// todo: docs
func (f *File) Close() error {
	return f.ref.Close()
}
