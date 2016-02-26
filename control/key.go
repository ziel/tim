package control

import (
	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/control/errors"
)

// todo: docs
func (s *state) key(event termbox.Event) error {
	switch event.Key {
	case termbox.KeyCtrlC:
		return errors.Quit

	case termbox.KeyArrowUp:
		// todo: s.view.scrollUp

	case termbox.KeyArrowDown:
		// todo: s.view.scrollUp

	case termbox.KeyArrowLeft:
		s.model.MergeLeft()

	case termbox.KeyArrowRight:
		s.model.MergeRight()
	}

	switch event.Ch {
	case 'q':
		return errors.Quit

	case 'k':
		// todo: s.view.scrollUp

	case 'j':
		// todo: s.view.scrollUp

	case 'h':
		s.model.MergeLeft()

	case 'l':
		s.model.MergeRight()
	}

	return nil
}
