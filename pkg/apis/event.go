package apis

import "time"

type Event struct {
	Name      string
	CreatedAt time.Time
	Payload   interface{}
}

func NewEvent(name string, payload interface{}) *Event {
	return &Event{Name: name, CreatedAt: time.Now(), Payload: payload}
}

type EventHandlerFunc func(event *Event)

type EventLoop interface {
	On(name string, handler EventHandlerFunc)
	Emit(event *Event)
}
