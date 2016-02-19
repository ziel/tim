// todo: docs
//
package control

import (
	"errors"
	"log"

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
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	go func() {
		for {
			ev := termbox.PollEvent()

			if ev.Type == termbox.EventInterrupt {
				log.Println("EXIT")
				return
			}

			events <- ev
		}
	}()

	for {
		select {
		case event := <-events:
			if err := handle(&event); err != nil {
				switch err {
				case errQuit:
					termbox.Interrupt()
					return nil
				}
				return err
			}
		}

		currentView.Draw()
	}

	return nil
}

// todo: docs
func handle(event *termbox.Event) error {
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
func key(event *termbox.Event) error {
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
