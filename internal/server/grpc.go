package server

import (
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcServer implements a gRPC Server for the Order service
type GrpcServer struct {
	Server   *grpc.Server
	listener net.Listener
	errCh    chan error
}

// NewGrpcServer is a convenience func to create a GrpcServer
func NewGrpcServer(port string, services ...func(grpc.ServiceRegistrar)) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)

	for _, s := range services {
		s(server)
	}

	// use reflection to register service on gRPC server.
	reflection.Register(server)

	return GrpcServer{
		Server:   server,
		listener: lis,
		errCh:    make(chan error, 1),
	}, nil
}

// Start starts the server in the background, pushing any error to the error channel
func (g GrpcServer) Start() {
	go func() {
		if err := g.Server.Serve(g.listener); err != nil {
			g.errCh <- err
		}
	}()
}

// Stop stops the gRPC server
func (g GrpcServer) Stop() {
	g.Server.GracefulStop()
}

// Error returns the server's error channel
func (g GrpcServer) Error() chan error {
	return g.errCh
}
