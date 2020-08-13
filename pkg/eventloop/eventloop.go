package eventloop

import (
	"fmt"

	"github.com/mehdy/Assistant/pkg/apis"
)

const EventsQueueSize = 1024

type EventLoop struct {
	handlers   map[string][]apis.EventHandlerFunc
	eventQueue chan *apis.Event
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		eventQueue: make(chan *apis.Event, EventsQueueSize),
	}
}

func (e *EventLoop) Emit(event *apis.Event) {
	e.eventQueue <- event
}

func (e *EventLoop) Run() {
	for event := range e.eventQueue {
		handlers, ok := e.handlers[event.Name]
		if !ok {
			fmt.Printf("No event handler is registered for %q\n", event.Name)

			continue
		}

		for _, handler := range handlers {
			go handler(event)
		}
	}
}

func (e *EventLoop) On(name string, handler apis.EventHandlerFunc) {
	if _, ok := e.handlers[name]; !ok {
		e.handlers[name] = []apis.EventHandlerFunc{}
	}

	e.handlers[name] = append(e.handlers[name], handler)
}
