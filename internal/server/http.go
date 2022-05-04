package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	orderpb "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
)

// HTTPServer implements a HTTP server for the Order Service
type HTTPServer struct {
	server       *http.Server
	orderService orderpb.OrderServiceServer // the same order service we injected into the gRPC server
	errCh        chan error
}

// NewHTTPServer is a convenience func to create a HTTPServer
func NewHTTPServer(port string, orderService orderpb.OrderServiceServer) HTTPServer {
	router := gin.Default()

	rs := HTTPServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
		errCh:        make(chan error, 1),
	}

	// register routes
	router.POST("/order", rs.create)
	router.GET("/order/:id", rs.retrieve)
	router.PUT("/order", rs.update)
	router.DELETE("/order", rs.delete)
	router.GET("/orders", rs.list)

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
	return r.server.Close()
}

// Error returns the server's error channel
func (r HTTPServer) Error() chan error {
	return r.errCh
}

// create is a handler func that creates an order from an order request (JSON body)
func (r HTTPServer) create(c *gin.Context) {
	var req orderpb.CreateOrderRequest

	// unmarshal the order request
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}

	// use the order service to create the order from the req
	resp, err := r.orderService.Create(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}

func (r HTTPServer) retrieve(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r HTTPServer) update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r HTTPServer) delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r HTTPServer) list(c *gin.Context) {
	for i := 1; i <= 7; i++ {
		log.Println(i)
		time.Sleep(time.Second)
	}
	orders := make([]*orderpb.Order, 3)
	for i := 0; i < 3; i++ {
		order := &orderpb.Order{OrderId: int64(i + 1)}
		orders[i] = order
	}
	resp := &orderpb.ListOrderResponse{Orders: orders}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending orders response")
	}
}
