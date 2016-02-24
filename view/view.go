// todo: docs
//
package view

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
)

// todo: docs
type View interface {

	// todo: docs
	Draw()

	// todo: docs
	Resize(termbox.Event) error
}

// todo: docs
func Factory(m model.Model) (View, error) {
	paths := m.FilePaths()

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
