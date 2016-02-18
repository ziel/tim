package model

// todo: docs
type Model struct {

	// todo: docs
	Files []string
}

func Fake2FileModel() *Model {
	return &Model{
		Files: []string{"one", "two"},
	}
}

func Fake3FileModel() *Model {
	return &Model{
		Files: []string{"one", "two", "three"},
	}
}
