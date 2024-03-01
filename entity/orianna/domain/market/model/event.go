package model

import (
	"fmt"
	"time"
)

// EventType is an enum that represents the type of event.
type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeEarningsCall
	EventTypeDividend
)

var eventTypeMap = map[EventType]string{
	EventTypeUnknown:      "unknown",
	EventTypeEarningsCall: "earnings_call",
	EventTypeDividend:     "dividend",
}

func (x *EventType) String() string {
	return eventTypeMap[*x]
}

// Event is an entity that represents an event.
type Event struct {
	ID         string    `json:"id"`
	Symbol     string    `json:"symbol,omitempty"`
	Type       EventType `json:"type,omitempty"`
	OccurredAt time.Time `json:"occurred_at,omitempty"`
}

// NewEvent is to create a new event.
func NewEvent(symbol string, typ EventType, occurredAt time.Time) *Event {
	return &Event{
		ID:         fmt.Sprintf("%s-%d-%d", symbol, typ, occurredAt.Unix()),
		Symbol:     symbol,
		Type:       typ,
		OccurredAt: occurredAt,
	}
}
