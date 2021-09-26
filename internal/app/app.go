package app

import (
	"context"

	"github.com/leighlin0511/grpc_template/internal/server"
	orderpb "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
)

const (
	grpcPort = "50051"
	restPort = "8080"
)

// App is a convenience wrapper for all things needed to start
// and shutdown the Order microservice
type App struct {
	restServer server.RestServer
	grpcServer server.GrpcServer
}

// start starts the REST and gRPC Servers in the background
func (a App) start() {
	a.restServer.Start() // non blocking now
	a.grpcServer.Start() // also non blocking :-)
}

// stop shuts down the servers
func (a App) shutdown() error {
	a.grpcServer.Stop()
	return a.restServer.Stop()
}

// newApp creates a new app with REST & gRPC servers
// this func performs all app related initialization
func newApp() (App, error) {
	orderService := orderpb.UnimplementedOrderServiceServer{}

	gs, err := server.NewGrpcServer(orderService, grpcPort)
	if err != nil {
		return App{}, err
	}

	return App{
		restServer: server.NewRestServer(orderService, restPort),
		grpcServer: gs,
	}, nil
}

// Run starts the app, handling any REST or gRPC server error
// and as well as app shutdown
func Run(ctx context.Context) error {
	app, err := newApp()
	if err != nil {
		return err
	}

	app.start()
	defer app.shutdown()

	select {
	case restErr := <-app.restServer.Error():
		return restErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-ctx.Done():
		return nil
	}
}
