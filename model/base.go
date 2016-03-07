package model

// todo: docs
type base struct {
	files map[string]*File
}

// todo: docs
func newBase(paths []string) (*base, error) {
	files := make(map[string]*File)

	for _, path := range paths {
		file, err := NewFile(path)

		if err != nil {
			return nil, err
		}

		files[path] = file
	}

	return &base{files: files}, nil
}

// todo: docs
func (m *base) Paths() []string {
	result := make([]string, 0, len(m.files))

	for path := range m.files {
		result = append(result, path)
	}

	return result
}

// todo: docs
func (m *base) File(path string) *File {
	return m.files[path]
}

// todo: docs
func (m *base) Close() error {
	for _, file := range m.files {
		err := file.Close()

		if err != nil {
			// don't bother trying the rest
			return err
		}
	}

	return nil
}
