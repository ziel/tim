// todo: docs
//
package model

import "github.com/ziel/tim/timerror"

// todo: docs
type Model interface {
	FilePaths() []string
	MergeRight() error
	MergeLeft() error
}

// todo: docs
func Factory(files []string) (Model, error) {
	switch len(files) {
	case 0, 1:
		return nil, timerror.TooFewFiles
	case 2:
		return NewTwoUp(files[0], files[1]), nil
	case 3:
		return NewThreeUp(files[0], files[1], files[2]), nil
	}

	return nil, timerror.TooManyFiles
}
