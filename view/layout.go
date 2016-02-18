package view

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/ziel/tim/model"
)

// note for later: can draw connectors like this:
// ui.SetRune(maxX/2, maxY/2+1, '─')

type Layout interface {

	// todo: docs
	Update(int, int)

	// todo: docs
	Handler(*gocui.Gui) error
}

// todo: docs (view description)
type viewDescr struct {

	// the view name
	Name string

	// the coordinates
	X0, Y0, X1, Y1 int
}

// todo: fancier
func layoutFactory(m *model.Model) Layout {
	switch m.Files.length {
	case 3:
		return newThreeUpLayout()
	case 2:
		return newTwoUpLayout()
	}

	// todo: fixup
	panic("I can only handle two or three files for now")
}

// todo: docs
func placeView(ui *gocui.Gui, ds viewDescr) error {
	view, err := ui.SetView(ds.Name, ds.X0, ds.Y0, ds.X1, ds.Y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		initView(view)
	}

	return nil
}

// todo: docs
func placeViews(ui *gocui.Gui, views ...viewDescr) error {
	for _, v := range views {
		if err := placeView(ui, v); err != nil {
			return err
		}
	}

	return nil
}

// todo: docs
func initView(view *gocui.View) {
	view.BgColor = gocui.ColorDefault
	view.FgColor = gocui.ColorWhite
	view.Frame = true

	fmt.Fprintln(view, "I am: ", view.Name())
}
