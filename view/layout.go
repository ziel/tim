package view

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// todo: docs
type Layout struct {
	// todo: docs
	nFlexibleElements int

	// todo: docs
	minWidth int

	// todo: docs
	elements []Element
}

func NewLayout(elements []Element) *Layout {
	result := &Layout{
		nFlexibleElements: 0,
		minWidth:          0,
		elements:          elements,
	}

	result.init()
	return result
}

func (l *Layout) init() {
	l.minWidth = 0
	l.nFlexibleElements = 0

	for _, e := range l.elements {
		min, max := e.WidthConstraints()
		l.minWidth += min

		if max <= 0 {
			l.nFlexibleElements += 1
		}
	}
}

// todo: docs
func widthForElement(width int, element Element) int {
	if _, max := element.WidthConstraints(); max > 0 {
		if max < width {
			return max
		}
	}
	return width
}

func (l *Layout) Update(width, height int) error {
	available := width - l.minWidth
	divided := available / l.nFlexibleElements

	x0 := 0

	for _, e := range l.elements {
		x1 := x0 + widthForElement(divided, e)

		if err := e.SetRect(NewRect(x0, 0, x1, height)); err != nil {
			return err
		}

		x0 = x1 + 1
	}

	return nil
}

func (l *Layout) Draw() {
	for _, el := range l.elements {
		el.Draw()
	}

	termbox.Flush()
}

func (l *Layout) Resize(ev termbox.Event) error {
	if ev.Type != termbox.EventResize {
		return fmt.Errorf("Resize called with non-resize event: %s", ev)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	if err := l.Update(ev.Width, ev.Height); err != nil {
		return err
	}

	l.Draw()
	return nil
}
