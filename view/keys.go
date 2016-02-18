package view

import (
	"log"

	"github.com/jroimartin/gocui"
)

// todo: consider a control package for this
func quit(ui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// todo: docs
func bindKeys(ui *gocui.Gui) {
	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := ui.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}
