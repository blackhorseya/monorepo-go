package wirex

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
)

// Injector is the injector interface.
type Injector struct {
	A *configx.Application
}
