package model

import "fmt"

// todo: docs
type ThreeUp struct {
	*base
}

// todo: docs
func newThreeUp(paths []string) (*ThreeUp, error) {
	if len(paths) != 3 {
		msg := "ThreeUp Model needs exactly 3 files. have: %d"
		return nil, fmt.Errorf(msg, len(paths))
	}

	base, err := newBase(paths)

	if err != nil {
		return nil, err
	}

	return &ThreeUp{base: base}, nil
}

// todo: docs
func (tu *ThreeUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *ThreeUp) MergeLeft() error {
	return nil
}
