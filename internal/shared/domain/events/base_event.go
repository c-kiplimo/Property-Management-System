package events

import (
	"fmt"
	"time"
)

type DomainEvent interface {
	GetEventID() string
	GetEventType() string
	GetOccurredOn() time.Time
	GetAggregateID() string
}

type BaseEvent struct {
	EventID     string    `json:"event_id"`
	EventType   string    `json:"event_type"`
	OccurredOn  time.Time `json:"occurred_on"`
	AggregateID string    `json:"aggregate_id"`
}

func NewBaseEvent(eventType, aggregateID string) BaseEvent {
	return BaseEvent{
		EventID:     generateEventID(),
		EventType:   eventType,
		OccurredOn:  time.Now(),
		AggregateID: aggregateID,
	}
}

func (e BaseEvent) GetEventID() string       { return e.EventID }
func (e BaseEvent) GetEventType() string     { return e.EventType }
func (e BaseEvent) GetOccurredOn() time.Time { return e.OccurredOn }
func (e BaseEvent) GetAggregateID() string   { return e.AggregateID }

func generateEventID() string {
	// In a real implementation, use UUID
	return fmt.Sprintf("evt_%d", time.Now().UnixNano())
}
