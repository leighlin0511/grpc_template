package order

import (
	orderproto "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/app"
)

type IDatabase interface {
	Create(ID string, items []*orderproto.Item) (*orderproto.Order, error)
	Retrieve(ID string) (*orderproto.Order, error)
	Update(o *orderproto.Order) (*orderproto.Order, error)
	Delete(ID string) (string, error)
	List() ([]*orderproto.Order, error)
}

func NewMemoryDB() *MemDatabase {
	return &MemDatabase{
		Orders: make(map[string]*orderproto.Order),
	}
}
