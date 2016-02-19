package model

import "fmt"

// todo: docs
type Model interface {
	FilePaths() []string
}

type Two struct {
	filePaths [2]string
}

func newTwo(path1, path2 string) *Two {
	return &Two{
		filePaths: [2]string{path1, path2},
	}
}

func (dm *Two) FilePaths() []string {
	return dm.filePaths[:]
}

type Three struct {
	filePaths [3]string
}

func newThree(path1, path2, path3 string) *Three {
	return &Three{
		filePaths: [3]string{path1, path2, path3},
	}
}

func (tm *Three) FilePaths() []string {
	return tm.filePaths[:]
}

func Factory(files []string) (Model, error) {

	switch len(files) {
	case 0, 1:
		return nil, fmt.Errorf("%s\n", "I need at least 2 files to compare.")
	case 2:
		return newTwo(files[0], files[1]), nil
	case 3:
		return newThree(files[0], files[1], files[2]), nil
	}

	return nil, fmt.Errorf("%s\n", "I can't compare more than 3 files.")
}
