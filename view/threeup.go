package view

import "github.com/jroimartin/gocui"

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

func (l *threeUpLayout) Handler(ui *gocui.Gui) error {
	l.Update(ui.Size())
	return placeViews(ui, l.left, l.mid, l.right)
}
