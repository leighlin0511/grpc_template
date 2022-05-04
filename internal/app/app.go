package app

import (
	"context"
	"strconv"
	"time"

	"github.com/leighlin0511/grpc_template/internal/app/config"
	"github.com/leighlin0511/grpc_template/internal/server"
	"github.com/leighlin0511/grpc_template/pkg/service"
)

// App is a convenience wrapper for all things needed to start
// and shutdown the Order microservice
type App struct {
	httpServer   server.HTTPServer
	grpcServer   server.GrpcServer
	shutdownChan <-chan struct{}
}

// start starts the REST and gRPC Servers in the background
func (a App) start() {
	a.httpServer.Start() // non blocking now
	a.grpcServer.Start() // also non blocking :-)
}

// stop shuts down the servers
func (a App) shutdown() error {
	a.grpcServer.Stop()
	return a.httpServer.Stop()
}

// newApp creates a new app with REST & gRPC servers
// this func performs all app related initialization
func newApp(conf *config.Configuration) (App, error) {
	ctx := context.Background()

	orderService, err := service.NewOrderService(conf)
	if err != nil {
		return App{}, err
	}

	gs, err := server.NewGrpcServer(
		strconv.Itoa(conf.Server.GrpcPort),
		orderService.RegisterGRPC,
	)
	if err != nil {
		return App{}, err
	}
	wait := server.GracefulShutdown(ctx, conf.Server.ShutdownTimeout, map[string]server.Operation{
		"operation1": shutdownOperation1,
		"operation2": shutdownOperation2,
		"operation3": shutdownOperation3,
	})
	return App{
		httpServer: server.NewHTTPServer(
			strconv.Itoa(conf.Server.HTTPPort),
			orderService,
		),
		grpcServer:   gs,
		shutdownChan: wait,
	}, nil
}

// Run starts the app, handling any REST or gRPC server error
// and as well as app shutdown
func Run(conf *config.Configuration) error {
	app, err := newApp(conf)
	if err != nil {
		return err
	}

	app.start()
	defer app.shutdown()

	select {
	case httpErr := <-app.httpServer.Error():
		return httpErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-app.shutdownChan:
		return nil
	}
}

func shutdownOperation1() error {
	// mock expensive operation
	time.Sleep(4 * time.Second)
	return nil
}

func shutdownOperation2() error {
	// mock expensive operation
	time.Sleep(5 * time.Second)
	return nil
}

func shutdownOperation3() error {
	// mock expensive operation
	time.Sleep(6 * time.Second)
	return nil
}

// reference
// https://levelup.gitconnected.com/the-golang-microservice-toolkit-7521516ee4b
