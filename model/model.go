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
		return newTwoUp(files)
	case 3:
		return newThreeUp(files)
	}

	return nil, timerror.TooManyFiles
}
