package view

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/ziel/tim/model"
)

func createGui() *gocui.Gui {
	result := gocui.NewGui()

	if err := result.Init(); err != nil {
		panic(err)
	}

	result.BgColor = gocui.ColorDefault
	result.FgColor = gocui.ColorCyan

	return result
}

func Display(m *model.Model) {
	ui := createGui()
	defer ui.Close()

	layout := layoutFactory(m)
	ui.SetLayout(layout.Handler)

	bindKeys(ui)
	err := ui.MainLoop()

	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
