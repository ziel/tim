package model

// todo: docs
type ThreeUp struct {
	files [3]*File
}

// todo: docs
func NewThreeUp(path1, path2, path3 string) *ThreeUp {
	return nil
}

// todo: docs
func (tu *ThreeUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *ThreeUp) MergeLeft() error {
	return nil
}
