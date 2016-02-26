package view

import (
	"fmt"

	"github.com/ziel/tim/view/errors"
)

// todo: docs
// Implements Element
type Connector struct {
	rect Rect
}

// Width needed to draw a connector.
//
const ConnectorWidth = 3

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
func (c *Connector) Draw() {
	// todo
}
