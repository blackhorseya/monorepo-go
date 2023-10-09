package contextx

import (
	"context"

	"go.uber.org/zap"
)

// Contextx extends google's context to support logging methods
type Contextx struct {
	context.Context
	*zap.Logger
}
