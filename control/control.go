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
type controller struct {
	model     model.Model
	view      view.View
	waitGroup sync.WaitGroup
	cleaners  []func()
}

// todo: docs
var instance *controller

// todo: docs
func newController() *controller {
	return &controller{
		cleaners: make([]func(), 0),
	}
}

// todo: docs
func Init(paths []string) error {
	instance = newController()
	defer instance.cleanup()

	if err := instance.initModel(paths); err != nil {
		return err
	}

	if err := instance.initView(paths); err != nil {
		return err
	}

	if err := instance.initTermbox(); err != nil {
		return err
	}

	return instance.loop()
}

// todo: docs
func (c *controller) addCleaner(fn func()) {
	c.cleaners = append(c.cleaners, fn)
}

// todo: docs
func (c *controller) cleanup() {
	for _, fn := range c.cleaners {
		fn()
	}
}

// todo: docs
func (c *controller) initModel(paths []string) error {
	mdl, err := model.Factory(paths)
	if err != nil {
		return err
	}

	c.addCleaner(func() {
		if err := mdl.Close(); err != nil {
			log.Println(err)
		}
	})

	c.model = mdl
	return nil
}

// todo: docs
func (c *controller) initView(paths []string) error {
	viw, err := view.Factory(paths)

	if err != nil {
		return err
	}

	c.view = viw
	return nil
}

// todo: docs
func (c *controller) initTermbox() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	c.addCleaner(termbox.Close)

	termbox.SetInputMode(termbox.InputAlt)
	termbox.Flush()

	w, h := termbox.Size()
	return c.view.Update(w, h)
}

// todo: docs
func (c *controller) loop() error {
	events := c.eventProducer()
	result := c.eventConsumer(events)

	c.waitGroup.Wait()
	return <-result
}

// todo: docs
func (c *controller) activateGoroutine(fn func()) {
	c.waitGroup.Add(1)

	go func() {
		defer c.waitGroup.Done()
		fn()
	}()
}

// todo: docs
func (c *controller) eventProducer() <-chan termbox.Event {
	const bufsize int = 10
	events := make(chan termbox.Event, bufsize)

	c.activateGoroutine(func() {
		for {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventInterrupt {
				return
			}
			events <- ev
		}
	})

	return events
}

// todo: docs
func (c *controller) eventConsumer(events <-chan termbox.Event) <-chan error {
	result := make(chan error, 1)

	c.activateGoroutine(func() {
		for {
			c.view.Draw()

			if err := c.handle(<-events); err != nil {
				if err != errors.Quit {
					result <- err
				}

				close(result)
				termbox.Interrupt()

				return
			}
		}
	})

	return result
}

// todo: docs
func (c *controller) handle(event termbox.Event) error {
	switch event.Type {
	case termbox.EventKey:
		return c.key(event)

	case termbox.EventResize:
		return c.view.Resize(event)

	case termbox.EventError:
		return event.Err
	}

	return nil
}
