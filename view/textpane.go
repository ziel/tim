package view

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// todo: docs
// Implements Element
type TextPane struct {
	rect  Rect
	title string
}

// todo: docs
// todo: fix right margin issue
const TextPaneMinWidth = 5

func NewTextPane(title string) *TextPane {
	padded := fmt.Sprintf(" %s", title)
	return &TextPane{title: padded}
}

// Implementation for Element
//
func (t *TextPane) SetRect(rect *Rect) error {
	t.rect = *rect
	return nil
}

// Implementation for Element
// TextPane is flexible
//
func (t *TextPane) WidthConstraints() (int, int) {
	return TextPaneMinWidth, 0
}

// Implementation for Element
//
func (t *TextPane) Draw() {
	t.drawTitle(t.rect.Y0)
	t.drawBorders()

	// todo: draw doc
}

// todo: docs
func (t *TextPane) drawTitle(y int) {
	fg, bg := textColor()
	clearRegion(t.rect.X0, t.rect.X1, y, y)

	for i, chr := range []rune(t.title) {
		current := t.rect.X0 + i

		if current > t.rect.X1 {
			break
		}

		termbox.SetCell(current, y, chr, fg, bg)
	}
}

// todo: docs
// todo: move double connector height into general view stuff
//
func (t *TextPane) drawBorders() {
	lineDblHorizontal(t.rect.Y0+1, t.rect.X0, t.rect.X1)
}
