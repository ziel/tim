package view

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
)

type twoUpLayout struct {
	left  viewDescr
	right viewDescr
}

func newTwoUpLayout() *twoUpLayout {
	return &twoUpLayout{
		left:  viewDescr{Name: "left"},
		right: viewDescr{Name: "right"},
	}
}

type twoUp struct {
	model *model.Two
}

// todo: docs
func (u2 *twoUp) Draw() {
	// todo: impl redraw
	termbox.Flush()
}

// todo: docs
func (u2 *twoUp) Resize(event *termbox.Event) error {
	// todo: impl
	return nil
}

// todo: docs
func (l *twoUpLayout) Update(maxX, maxY int) {
	l.left.X0 = 0
	l.left.Y0 = 0
	l.left.X1 = maxX/2 - 1
	l.left.Y1 = maxY - 1

	l.right.X0 = maxX/2 + 1
	l.right.Y0 = 0
	l.right.X1 = maxX - 1
	l.right.Y1 = maxY - 1
}
