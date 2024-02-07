package model

import (
	"encoding/json"
	"time"
)

func (x *MarketInfo) MarshalJSON() ([]byte, error) {
	type Alias MarketInfo

	return json.Marshal(&struct {
		*Alias
		QueriedAt string `json:"queried_at,omitempty"`
	}{
		Alias:     (*Alias)(x),
		QueriedAt: x.QueriedAt.AsTime().UTC().Format(time.RFC3339),
	})
}
