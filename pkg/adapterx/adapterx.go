package adapterx

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
}

// Grpc is the interface that wraps the grpc api method.
type Grpc interface {
	Servicer
}