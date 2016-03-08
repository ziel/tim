package view

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/view/errors"
)

// todo: docs
// Implements Element
type Connector struct {
	rect Rect
}

// Width needed to draw a connector.
//
const ConnectorWidth = 1

func NewConnector() *Connector {
	return &Connector{}
}

// Implementation for Element
//
func (c *Connector) SetRect(rect *Rect) error {
	if rect.Width() < ConnectorWidth {
		msg := fmt.Sprintf(
			"connector width must be %d, got %d",
			ConnectorWidth,
			rect.Width())

		return errors.LayoutError(msg)
	}

	c.rect = *rect
	return nil
}

// Implementation for Element
//
func (c *Connector) WidthConstraints() (int, int) {
	return ConnectorWidth, ConnectorWidth
}

// Implementation for Element
//
// todo: move double connector height into general view stuff
// todo: add individual connector drawing
//
func (c *Connector) Draw() {
	lineVertical(c.rect.X0, c.rect.Y0, c.rect.Y1)

	fg, bg := lineColor()
	termbox.SetCell(c.rect.X0, c.rect.Y0+1, '╪', fg, bg)
}
