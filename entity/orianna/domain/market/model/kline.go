package model

import (
	"time"
)

// KLine is a value object that represents a KLine data.
type KLine struct {
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`

	Interval Interval `json:"interval"`

	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`

	Closed bool `json:"closed"`
}
