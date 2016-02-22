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

var currentView view.View

// todo: docs
func Init(m model.Model) error {
	var err error
	currentView, err = view.Factory(m)

	if err != nil {
		return err
	}

	startTermbox()
	return nil
}

// todo: docs
var errQuit = errors.New("Quit")

// todo: docs
func startTermbox() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)

	currentView.Draw()

	if err := eventloop(); err != nil {
		panic(err)
	}
}

// todo: docs
func eventloop() error {
	var wg sync.WaitGroup
	wg.Add(2)

	events := eventProducer(wg.Done)
	result := eventConsumer(wg.Done, events)

	wg.Wait()
	return <-result
}

// todo: docs
func eventProducer(done func()) <-chan termbox.Event {
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	go func() {
		defer done()

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
func eventConsumer(done func(), events <-chan termbox.Event) <-chan error {
	result := make(chan error)

	go func() {
		defer done()

		for {
			currentView.Draw()
			err := handle(<-events)

			if err == nil {
				continue
			}

			switch err {
			case errQuit:
				close(result)
				termbox.Interrupt()

			default:
				result <- err
			}

			return
		}
	}()

	return result
}

// todo: docs
func handle(event termbox.Event) error {
	switch event.Type {
	case termbox.EventKey:
		return key(event)

	case termbox.EventResize:
		return currentView.Resize(event)

	case termbox.EventError:
		return event.Err
	}

	return nil
}

// todo: docs
func key(event termbox.Event) error {
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
