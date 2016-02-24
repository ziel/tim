package view

import "fmt"

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
		es := "Connector width must be %d, got %d"
		return fmt.Errorf(es, ConnectorWidth, rect.Width())
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
