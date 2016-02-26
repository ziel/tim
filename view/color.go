package view

import "github.com/nsf/termbox-go"

// todo: docs
// todo: configurable later
func lineColor() (termbox.Attribute, termbox.Attribute) {
	return termbox.ColorBlack, termbox.ColorDefault
}

// todo: docs
// todo: configurable later
func textColor() (termbox.Attribute, termbox.Attribute) {
	return termbox.ColorCyan, termbox.ColorDefault
}
