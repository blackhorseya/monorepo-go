package agg

import (
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/model"
)

// Asset is an aggregate root.
type Asset struct {
	model.Car
}
