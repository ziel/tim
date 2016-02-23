package model

// todo: docs
type ThreeUp struct {
	filePaths [3]string
}

// todo: docs
func NewThreeUp(path1, path2, path3 string) *ThreeUp {
	return &ThreeUp{
		filePaths: [3]string{path1, path2, path3},
	}
}

// todo: docs
func (tu *ThreeUp) FilePaths() []string {
	return tu.filePaths[:]
}

// todo: docs
func (tu *ThreeUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *ThreeUp) MergeLeft() error {
	return nil
}
