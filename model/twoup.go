package model

// todo: docs
type TwoUp struct {
	filePaths [2]string
}

// todo: docs
func NewTwoUp(path1, path2 string) *TwoUp {
	return &TwoUp{
		filePaths: [2]string{path1, path2},
	}
}

// todo: docs
func (tu *TwoUp) FilePaths() []string {
	return tu.filePaths[:]
}

// todo: docs
func (tu *TwoUp) MergeRight() error {
	return nil
}

// todo: docs
func (tu *TwoUp) MergeLeft() error {
	return nil
}
