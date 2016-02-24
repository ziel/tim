// todo: docs
//
package control

import (
	"log"
	"sync"

	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/model"
	"github.com/ziel/tim/timerror"
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

	defer func() {
		if err := cModel.Close(); err != nil {
			log.Println(err)
		}
	}()

	cView, err := view.Factory(paths)
	if err != nil {
		return err
	}

	controller = &state{
		model: cModel,
		view:  cView,
	}

	return controller.loop()
}

// todo: docs
func (s *state) loop() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)

	events := s.eventProducer()
	result := s.eventConsumer(events)

	s.wg.Wait()
	return <-result
}

// todo: docs
func (s *state) eventProducer() <-chan termbox.Event {
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	s.wg.Add(1)

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

	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		for {
			s.view.Draw()

			if err := s.handle(<-events); err != nil {
				switch err {
				case timerror.Quit:
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
