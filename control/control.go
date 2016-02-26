// todo: docs
//
package control

import (
	"log"
	"sync"

	"github.com/nsf/termbox-go"
	"github.com/ziel/tim/control/errors"
	"github.com/ziel/tim/model"
	"github.com/ziel/tim/view"
)

// todo: docs
type state struct {
	model     model.Model
	view      view.View
	waitGroup sync.WaitGroup
	cleaners  []func()
}

// todo: docs
var controller *state

// todo: docs
func newController() *state {
	return &state{
		cleaners: make([]func(), 0),
	}
}

// todo: docs
func Init(paths []string) error {
	controller = newController()
	defer controller.cleanup()

	if err := controller.initModel(paths); err != nil {
		return err
	}

	if err := controller.initView(paths); err != nil {
		return err
	}

	if err := controller.initTermbox(); err != nil {
		return err
	}

	return controller.loop()
}

// todo: docs
func (s *state) initModel(paths []string) error {
	mdl, err := model.Factory(paths)
	if err != nil {
		return err
	}

	s.addCleaner(func() {
		if err := mdl.Close(); err != nil {
			log.Println(err)
		}
	})

	s.model = mdl
	return nil
}

// todo: docs
func (s *state) initView(paths []string) error {
	viw, err := view.Factory(paths)

	if err != nil {
		return err
	}

	s.view = viw
	return nil
}

// todo: docs
func (s *state) addCleaner(fn func()) {
	s.cleaners = append(s.cleaners, fn)
}

// todo: docs
func (s *state) cleanup() {
	for _, fn := range s.cleaners {
		fn()
	}
}

// todo: docs
func (s *state) loop() error {
	events := s.eventProducer()
	result := s.eventConsumer(events)

	s.waitGroup.Wait()
	return <-result
}

// todo: docs
func (s *state) initTermbox() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	s.addCleaner(termbox.Close)

	termbox.SetInputMode(termbox.InputAlt)
	termbox.Flush()

	w, h := termbox.Size()
	return s.view.Update(w, h)
}

// todo: docs
func (s *state) eventProducer() <-chan termbox.Event {
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	s.waitGroup.Add(1)

	go func() {
		defer s.waitGroup.Done()

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
	result := make(chan error, 1)

	s.waitGroup.Add(1)

	go func() {
		defer s.waitGroup.Done()

		for {
			s.view.Draw()

			if err := s.handle(<-events); err != nil {
				if err != errors.Quit {
					result <- err
				}

				close(result)
				termbox.Interrupt()

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
