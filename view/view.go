// todo: docs
//
package view

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
	"github.com/ziel/tim/view/errors"
)

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
func Factory(m model.Model) (View, error) {
	if m == nil {
		return nil, errors.NilModelError
	}

	paths := m.Paths()
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
