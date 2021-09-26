package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	orderpb "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
)

// RestServer implements a REST server for the Order Service
type RestServer struct {
	server       *http.Server
	orderService orderpb.OrderServiceServer // the same order service we injected into the gRPC server
}

// NewRestServer is a convenience func to create a RestServer
func NewRestServer(orderService orderpb.OrderServiceServer, port string) RestServer {
	router := gin.Default()

	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
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
func (r RestServer) Start() error {
	return r.server.ListenAndServe()
}

// create is a handler func that creates an order from an order request (JSON body)
func (r RestServer) create(c *gin.Context) {
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

func (r RestServer) retrieve(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) list(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
