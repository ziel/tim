package view

import "github.com/nsf/termbox-go"

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
	return &Layout{
		nFlexibleElements: 0,
		minWidth:          0,
		elements:          elements,
	}
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
		return max
	}
	return width
}

func (l *Layout) Update(width, height int) {
	available := width - l.minWidth
	divided := available / l.nFlexibleElements

	x0 := 0

	for _, e := range l.elements {
		x1 := x0 + widthForElement(divided, e)
		if _, max := e.WidthConstraints(); max > 0 {
			x1 = x0 + max
		} else {
			x1 = x0 + divided
		}

		x0 = x1
	}

	//todo: update contained elements
	// start with minimum layout
	// divide available width among flexible elements
	// expanding to fill
}

func (l *Layout) Draw() {
	// todo: implement
}

func (l *Layout) Resize(ev termbox.Event) error {
	// todo: implement
	return nil
}
