package model

import "fmt"

// todo: docs
type TwoUp struct {
	*base
}

// todo: docs
func newTwoUp(paths []string) (*TwoUp, error) {
	if len(paths) != 2 {
		msg := "TwoUp Model needs exactly 3 files. have: %d"
		return nil, fmt.Errorf(msg, len(paths))
	}

	base, err := newBase(paths)

	if err != nil {
		return nil, err
	}

	return &TwoUp{base: base}, nil
}

// todo: docs
func (tu *TwoUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *TwoUp) MergeLeft() error {
	return nil
}
