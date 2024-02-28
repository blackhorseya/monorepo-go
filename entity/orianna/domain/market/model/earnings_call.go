package model

import (
	"time"

	"github.com/google/uuid"
)

// EarningsCall is an entity that represents an earnings call
type EarningsCall struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Symbol     string    `json:"symbol,omitempty"`
	Host       string    `json:"host,omitempty"`
	OccurredAt time.Time `json:"occurred_at,omitempty"`
}
