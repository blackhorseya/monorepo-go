package agg

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/strategy/model"
)

// Strategy is an aggregate root.
type Strategy struct {
	*model.Strategy `json:",inline"`
	Targets         []*model.Target `json:"targets"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
