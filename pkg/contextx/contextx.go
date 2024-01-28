package contextx

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// KeyCtx is the key of contextx.
const KeyCtx = "contextx"

// Contextx extends google's context to support logging methods.
type Contextx struct {
	context.Context
	*zap.Logger
}

// Background returns a non-nil, empty Contextx. It is never canceled, has no values, and has no deadline.
func Background() Contextx {
	return Contextx{
		Context: context.Background(),
		Logger:  zap.L(),
	}
}

// WithLogger returns a copy of parent in which the value associated with key is val.
func WithLogger(logger *zap.Logger) Contextx {
	return Contextx{
		Context: context.Background(),
		Logger:  logger,
	}
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent Contextx, key, val interface{}) Contextx {
	return Contextx{
		Context: context.WithValue(parent.Context, key, val),
		Logger:  parent.Logger,
	}
}

// WithTimeout returns a copy of the parent context with the timeout adjusted to be no later than d.
func WithTimeout(parent Contextx, d time.Duration) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent.Context, d)

	return Contextx{
		Context: ctx,
		Logger:  parent.Logger,
	}, cancel
}

// WithCancel returns a copy of the parent context with a new Done channel.
func WithCancel(parent Contextx) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent.Context)

	return Contextx{
		Context: ctx,
		Logger:  parent.Logger,
	}, cancel
}
