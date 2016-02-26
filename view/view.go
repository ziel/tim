// todo: docs
//
package view

import "github.com/nsf/termbox-go"

// todo: docs
type View interface {

	// todo: docs
	Draw()

	// todo: docs
	Update(width, height int) error

	// todo: docs
	Resize(termbox.Event) error
}

// todo: docs
func Factory(paths []string) (View, error) {
	npaths := len(paths)
	nelems := (npaths * 2) - 1
	elems := make([]Element, 0, nelems)

	for i, path := range paths {
		elems = append(elems, NewTextPane(path))

		if i < (npaths - 1) {
			elems = append(elems, NewConnector())
		}
	}

	return NewLayout(elems), nil
}
