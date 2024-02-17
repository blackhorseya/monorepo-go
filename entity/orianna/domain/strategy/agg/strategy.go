package agg

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/strategy/model"
)

// Strategy is an aggregate root.
type Strategy struct {
	*model.Strategy `json:",inline"`
}
