package model

import (
	"time"
)

// Strategy is an entity.
type Strategy struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// Target is a value object.
type Target struct {
	Symbol     string    `json:"symbol"`
	OccurredAt time.Time `json:"occurred_at"`
}
