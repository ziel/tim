package model

// todo: docs
type TwoUp struct {
	files [2]*File
}

// todo: docs
func NewTwoUp(path1, path2 string) *TwoUp {
	return nil
}

// todo: docs
func (tu *TwoUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *TwoUp) MergeLeft() error {
	return nil
}
