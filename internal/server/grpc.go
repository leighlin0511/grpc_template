package server

import (
	"net"

	orderpb "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
	"google.golang.org/grpc"
)

// GrpcServer implements a gRPC Server for the Order service
type GrpcServer struct {
	server   *grpc.Server
	listener net.Listener
	errCh    chan error
}

// NewGrpcServer is a convenience func to create a GrpcServer
func NewGrpcServer(service orderpb.OrderServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(server, service)

	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error, 1),
	}, nil
}

// Start starts the server in the background, pushing any error to the error channel
func (g GrpcServer) Start() {
	go func() {
		if err := g.server.Serve(g.listener); err != nil {
			g.errCh <- err
		}
	}()
}

// Stop stops the gRPC server
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

// Error returns the server's error channel
func (g GrpcServer) Error() chan error {
	return g.errCh
}
