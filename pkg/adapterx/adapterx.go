package adapterx

import (
	"github.com/gin-gonic/gin"
)

// Servicer is the interface that wraps the basic Serve method.
type Servicer interface {
	// Start a service asynchronously.
	Start() error

	// AwaitSignal waits for a signal to shutdown the service.
	AwaitSignal() error
}

// Restful is the interface that wraps the restful api method.
type Restful interface {
	Servicer

	// InitRouting init the routing of restful api.
	InitRouting() error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}

// Grpc is the interface that wraps the grpc api method.
type Grpc interface {
	Servicer
}

// Cronjob is the interface that wraps the cronjob method.
type Cronjob interface {
	Servicer
}
