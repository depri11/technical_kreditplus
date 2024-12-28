package interfaces

import (
	"context"

	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type CustomerRepository interface {
	GetCustomer(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error)
	CreateCusstomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error
	UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, id string) error
}
