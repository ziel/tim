package control

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/control/errors"
)

// todo: docs
func (c *controller) key(event termbox.Event) error {
	switch event.Key {
	case termbox.KeyCtrlC:
		return errors.Quit

	case termbox.KeyArrowUp:
		// todo: c.view.scrollUp

	case termbox.KeyArrowDown:
		// todo: c.view.scrollUp

	case termbox.KeyArrowLeft:
		c.model.MergeLeft()

	case termbox.KeyArrowRight:
		c.model.MergeRight()
	}

	switch event.Ch {
	case 'q':
		return errors.Quit

	case 'k':
		// todo: c.view.scrollUp

	case 'j':
		// todo: c.view.scrollUp

	case 'h':
		c.model.MergeLeft()

	case 'l':
		c.model.MergeRight()
	}

	return nil
}
