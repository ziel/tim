package view

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
)

type threeUpLayout struct {
	left  viewDescr
	mid   viewDescr
	right viewDescr
}

func newThreeUpLayout() *threeUpLayout {
	return &threeUpLayout{
		left:  viewDescr{Name: "left"},
		mid:   viewDescr{Name: "mid"},
		right: viewDescr{Name: "right"},
	}
}

type threeUp struct {
	model *model.Three
}

// todo: docs
func (u3 *threeUp) Draw() {
	// todo: impl redraw
	termbox.Flush()
}

// todo: docs
func (u3 *threeUp) Resize(event termbox.Event) error {
	// todo: impl
	return nil
}

// todo: docs
func (l *threeUpLayout) Update(maxX, maxY int) {
	l.left.X0 = 0
	l.left.Y0 = 0
	l.left.X1 = maxX/3 - 1
	l.left.Y1 = maxY - 1

	l.mid.X0 = maxX/3 + 1
	l.mid.Y0 = 0
	l.mid.X1 = maxX*2/3 - 1
	l.mid.Y1 = maxY - 1

	l.right.X0 = maxX*2/3 + 1
	l.right.Y0 = 0
	l.right.X1 = maxX - 1
	l.right.Y1 = maxY - 1
}
