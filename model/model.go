// todo: docs
//
package model

import "github.com/ziel/tim/timerror"

// todo: docs
type Model interface {

	// todo: docs
	Paths() []string

	// todo: docs
	File(path string) *File

	// todo: docs
	MergeRight() error

	// todo: docs
	MergeLeft() error

	// todo: docs
	Close() error
}

// todo: docs
func Factory(files []string) (Model, error) {
	switch len(files) {
	case 0, 1:
		return nil, timerror.TooFewFiles
	case 2:
		return newFileModel(files)
	case 3:
		return newFileModel(files)
	}

	return nil, timerror.TooManyFiles
}

// todo: docs
type fileModel struct {
	files map[string]*File
}

// todo: docs
func newFileModel(paths []string) (*fileModel, error) {
	files := make(map[string]*File)

	for _, path := range paths {
		file, err := NewFile(path)

		if err != nil {
			return nil, err
		}

		files[path] = file
	}

	return &fileModel{files: files}, nil
}

// todo: docs
func (m *fileModel) Paths() []string {
	result := make([]string, len(m.files))

	for path := range m.files {
		result = append(result, path)
	}

	return result
}

// todo: docs
func (m *fileModel) File(path string) *File {
	return m.files[path]
}

// todo: docs
func (m *fileModel) MergeRight() error {
	return nil // todo
}

// todo: docs
func (m *fileModel) MergeLeft() error {
	return nil // todo
}

// todo: docs
func (m *fileModel) Close() error {
	for _, file := range m.files {
		err := file.Close()

		if err != nil {
			// don't bother trying the rest
			return err
		}
	}

	return nil
}
