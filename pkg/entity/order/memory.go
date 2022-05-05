package order

import (
	"errors"

	orderproto "github.com/leighlin0511/grpc_template/protobuf/generated/pkg/service/v1/order"
)

type MemDatabase struct {
	Orders map[string]*orderproto.Order
}

func (db *MemDatabase) Create(ID string, items []*orderproto.Item) (*orderproto.Order, error) {
	if _, ok := db.Orders[ID]; ok {
		return nil, errors.New("already exists")
	}
	var t float32
	for _, i := range items {
		t += i.Price
	}
	o := &orderproto.Order{
		OrderId: ID,
		Items:   items,
		Total:   t,
		Status:  orderproto.Order_PENDING,
	}
	db.Orders[ID] = o
	return o, nil
}

func (db *MemDatabase) Retrieve(ID string) (*orderproto.Order, error) {
	if o, ok := db.Orders[ID]; ok {
		return o, nil
	}
	return nil, errors.New("not found")
}

func (db *MemDatabase) Update(o *orderproto.Order) (*orderproto.Order, error) {
	if _, ok := db.Orders[o.OrderId]; !ok {
		return nil, errors.New("not found")
	}
	db.Orders[o.OrderId] = o
	return o, nil
}

func (db *MemDatabase) Delete(ID string) (string, error) {
	if o, ok := db.Orders[ID]; !ok {
		return "", errors.New("not found")
	} else {
		db.Orders[o.OrderId] = nil
	}
	return "successfully deleted", nil
}

func (db *MemDatabase) List() ([]*orderproto.Order, error) {
	result := make([]*orderproto.Order, 0)
	for _, o := range db.Orders {
		result = append(result, o)
	}
	if len(result) == 0 {
		return nil, errors.New("empty")
	}
	return result, nil
}
