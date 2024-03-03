package model

import (
	"time"
)

// Ticker is a value object that represents the ticker data.
type Ticker struct {
	Time   time.Time `json:"time,omitempty"`
	Volume float64   `json:"volume,omitempty"`
	Open   float64   `json:"open,omitempty"`
	High   float64   `json:"high,omitempty"`
	Close  float64   `json:"close,omitempty"`
	Low    float64   `json:"low,omitempty"`
}
