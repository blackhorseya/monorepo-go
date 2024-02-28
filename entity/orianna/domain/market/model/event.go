package model

import (
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
	Type       EventType `json:"type,omitempty"`
	OccurredAt time.Time `json:"occurred_at,omitempty"`
}
