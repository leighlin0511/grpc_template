package service

import (
	"context"
	"errors"

	"github.com/leighlin0511/grpc_template/internal/app/config"
	"github.com/leighlin0511/grpc_template/pkg/entity"
	orderentity "github.com/leighlin0511/grpc_template/pkg/entity/order"
	orderproto "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
	"google.golang.org/grpc"
)

type API struct {
	db orderentity.IDatabase
}
type SVC struct {
	orderproto.UnimplementedOrderServiceServer
	api API
}

func NewOrderService(conf *config.Configuration, db orderentity.IDatabase) (*SVC, error) {
	api := API{
		db: db,
	}
	return &SVC{
		api: api,
	}, nil
}

func (s *SVC) RegisterGRPC(registrar grpc.ServiceRegistrar) {
	orderproto.RegisterOrderServiceServer(registrar, s)
}

// creates a new order
func (s *SVC) Create(ctx context.Context, req *orderproto.CreateOrderRequest) (*orderproto.CreateOrderResponse, error) {
	o, err := s.api.db.Create(
		entity.NewID().String(),
		req.GetItems(),
	)
	if err != nil {
		return nil, errors.New("error creating")
	}
	return &orderproto.CreateOrderResponse{Order: o}, nil
}

// retrieves an existing order
func (s *SVC) Retrieve(context.Context, *orderproto.RetrieveOrderRequest) (*orderproto.RetrieveOrderResponse, error) {
	return nil, errors.New("not implemented")
}

// modifies an existing order
func (s *SVC) Update(context.Context, *orderproto.UpdateOrderRequest) (*orderproto.UpdateOrderResponse, error) {
	return nil, errors.New("not implemented")
}

// cancels an existing order
func (s *SVC) Delete(context.Context, *orderproto.DeleteOrderRequest) (*orderproto.DeleteOrderResponse, error) {
	return nil, errors.New("not implemented")
}

// lists existing orders
func (s *SVC) List(context.Context, *orderproto.ListOrderRequest) (*orderproto.ListOrderResponse, error) {
	return nil, errors.New("not implemented")
}
