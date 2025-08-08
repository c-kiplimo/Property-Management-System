package events

import (
	"log"
	"sync"
)

type EventHandler interface {
	Handle(event DomainEvent) error
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Subscribe(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(event DomainEvent) {
	d.mu.RLock()
	handlers := d.handlers[event.GetEventType()]
	d.mu.RUnlock()

	for _, handler := range handlers {
		go func(h EventHandler) {
			if err := h.Handle(event); err != nil {
				log.Printf("Error handling event %s: %v", event.GetEventType(), err)
			}
		}(handler)
	}
}
