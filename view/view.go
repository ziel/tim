// todo: docs
//
package view

import "github.com/nsf/termbox-go"

// todo: docs
type View interface {

	// todo: docs
	Draw()

	// todo: docs
	Resize(termbox.Event) error
}

// todo: docs
func Factory(paths []string) (View, error) {
	total := len(paths)
	elems := make([]Element, total*2-1)

	for i, path := range paths {
		elems = append(elems, NewTextPane(path))

		if i < total {
			elems = append(elems, NewConnector())
		}
	}

	return NewLayout(elems), nil
}
