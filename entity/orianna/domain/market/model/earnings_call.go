package model

import (
	"time"

	"github.com/google/uuid"
)

// EarningsCall is an entity that represents an earnings call
type EarningsCall struct {
	ID         uuid.UUID `json:"id"`
	Symbol     string    `json:"symbol"`
	Host       string    `json:"host"`
	OccurredAt time.Time `json:"occurred_at"`
}
