package view

type Element interface {

	// todo: docs
	SetRect(*Rect) error

	// todo: docs
	// returns min, max
	WidthConstraints() (int, int)

	// todo: docs
	Draw()
}
