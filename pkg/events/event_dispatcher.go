package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, &wg)
		}
		wg.Wait()
	}
	return nil
}

func (ev *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := ev.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ev.handlers[eventName] = append(ev.handlers[eventName], handler)
	return nil
}

func (ev *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if handlers, ok := ev.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ev *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := ev.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				ev.handlers[eventName] = append(ev.handlers[eventName][:i], ev.handlers[eventName][:i+1]...)
				return nil
			}
		}
	}
	return nil
}

func (ev *EventDispatcher) Clear() {
	ev.handlers = make(map[string][]EventHandlerInterface)
}
