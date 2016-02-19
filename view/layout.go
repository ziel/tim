package view

// note for later: can draw connectors like this:
// ui.SetRune(maxX/2, maxY/2+1, '─')

type Layout interface {

	// todo: docs
	Update(int, int)

	// todo: docs
	// Handler(*gocui.Gui) error
}

// todo: docs (view description)
type viewDescr struct {

	// the view name
	Name string

	// the coordinates
	X0, Y0, X1, Y1 int
}
