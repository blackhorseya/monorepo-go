package agg

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
)

// Event is an aggregate root that represents the event.
type Event struct {
	*model.Event
}
