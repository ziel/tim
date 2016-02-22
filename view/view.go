// todo: docs
//
package view

import (
	"fmt"

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
	switch v := m.(type) {
	case *model.Two:
		return &twoUp{model: v}, nil
	case *model.Three:
		return &threeUp{model: v}, nil
	}

	return nil, fmt.Errorf("cannot display model: %s", m)
}
