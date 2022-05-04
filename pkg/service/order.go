package service

import (
	"github.com/leighlin0511/grpc_template/internal/app/config"
	orderproto "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
	"google.golang.org/grpc"
)

type SVC struct {
	orderproto.UnimplementedOrderServiceServer
}

func NewOrderService(conf *config.Configuration) (*SVC, error) {
	return &SVC{}, nil
}

func (s *SVC) RegisterGRPC(registrar grpc.ServiceRegistrar) {
	orderproto.RegisterOrderServiceServer(registrar, s)
}
