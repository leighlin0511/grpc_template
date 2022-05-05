package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/leighlin0511/grpc_template/internal/app/config"
	orderpb "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const (
	contentType     = "content-type"
	grpcContentType = "application/grpc"
)

// HTTPServer implements a HTTP server for the Order Service
type HTTPServer struct {
	server       *http.Server
	orderService orderpb.OrderServiceServer // the same order service we injected into the gRPC server
	errCh        chan error
}

// NewHTTPServer is a convenience func to create a HTTPServer
func NewHTTPServer(ctx context.Context, orderService orderpb.OrderServiceServer, conf *config.Configuration, grpcHandler http.Handler) HTTPServer {
	gatewaymux := runtime.NewServeMux()
	portaddr := fmt.Sprintf(":%s", strconv.Itoa(conf.Server.HTTPPort))

	// register http service
	if err := orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, gatewaymux, portaddr, []grpc.DialOption{grpc.WithInsecure(), grpc.WithNoProxy()}); err != nil {
		log.Println("error when register order service http server")
	}

	mux := http.NewServeMux()
	mux.Handle("/template/", gatewaymux)
	rs := HTTPServer{
		server: &http.Server{
			Addr:    portaddr,
			Handler: grpcMessenger(ctx, &conf.Server, grpcHandler, mux),
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
		},
		orderService: orderService,
		errCh:        make(chan error, 1),
	}
	return rs
}

// Start starts the server
func (r HTTPServer) Start() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil {
			r.errCh <- err
		}
	}()
}

// Stop stops the server
func (r HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 300*time.Second)
	defer cancel()
	return r.server.Shutdown(ctx)
}

// Error returns the server's error channel
func (r HTTPServer) Error() chan error {
	return r.errCh
}

func grpcMessenger(ctx context.Context, conf *config.ServerConfig, grpcHandler http.Handler, httpHandler http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctxTimeout, cancel := context.WithTimeout(ctx, conf.Timeout)
		defer cancel()

		req := r.WithContext(ctxTimeout)
		contHeader := r.Header.Get(contentType)
		if r.ProtoMajor == 2 && strings.HasPrefix(contHeader, grpcContentType) {
			log.Printf("%s: %s, send to gRPC server", contentType, contHeader)
			grpcHandler.ServeHTTP(w, req)
		} else {
			log.Printf("%s: %s, send to HTTP server", contentType, contHeader)
			httpHandler.ServeHTTP(w, req)
		}
	}
	return h2c.NewHandler(http.HandlerFunc(f), &http2.Server{})
}
