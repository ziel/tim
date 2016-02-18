package view

import "github.com/jroimartin/gocui"

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

func (l *twoUpLayout) Handler(ui *gocui.Gui) error {
	l.Update(ui.Size())
	return placeViews(ui, l.left, l.right)
}
