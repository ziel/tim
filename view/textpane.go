package view

// todo: docs
// Implements Element
type TextPane struct {
	rect  Rect
	title string
}

// todo: docs
//
const TextPaneMinWidth = 1

func NewTextPane(title string) *TextPane {
	return &TextPane{title: title}
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
	// todo
}
