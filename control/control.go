// todo: docs
//
package control

import (
	"errors"
	"sync"

	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
	"github.com/ziel/tim/view"
)

// todo: docs
type state struct {
	model model.Model
	view  view.View
	wg    sync.WaitGroup
}

// todo: docs
var controller *state

// todo: docs
func Init(paths []string) error {
	cModel, err := model.Factory(paths)

	if err != nil {
		return err
	}

	cView, err := view.Factory(cModel)

	if err != nil {
		return err
	}

	controller = &state{
		model: cModel,
		view:  cView,
	}

	controller.init()
	return nil
}

// todo: docs
var errQuit = errors.New("Quit")

// todo: docs
func (s *state) init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)

	if err := s.eventLoop(); err != nil {
		panic(err)
	}
}

// todo: docs
func (s *state) eventLoop() error {
	s.wg.Add(2)

	events := s.eventProducer()
	result := s.eventConsumer(events)

	s.wg.Wait()
	return <-result
}

// todo: docs
func (s *state) eventProducer() <-chan termbox.Event {
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	go func() {
		defer s.wg.Done()

		for {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventInterrupt {
				return
			}
			events <- ev
		}
	}()

	return events
}

// todo: docs
func (s *state) eventConsumer(events <-chan termbox.Event) <-chan error {
	result := make(chan error)

	go func() {
		defer s.wg.Done()

		for {
			s.view.Draw()

			if err := s.handle(<-events); err != nil {
				switch err {
				case errQuit:
					close(result)
					termbox.Interrupt()

				default:
					result <- err
				}

				return
			}
		}
	}()

	return result
}

// todo: docs
func (s *state) handle(event termbox.Event) error {
	switch event.Type {
	case termbox.EventKey:
		return s.key(event)

	case termbox.EventResize:
		return s.view.Resize(event)

	case termbox.EventError:
		return event.Err
	}

	return nil
}

// todo: docs
func (s *state) key(event termbox.Event) error {
	switch event.Key {
	case termbox.KeyCtrlC:
		return errQuit
	}

	switch event.Ch {
	case 'q':
		return errQuit
	}

	return nil
}
