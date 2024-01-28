package wirex

import (
	"github.com/blackhorseya/monorepo-go/pkg/configx"
)

// Injector is the injector interface.
type Injector struct {
	A *configx.Application
}
