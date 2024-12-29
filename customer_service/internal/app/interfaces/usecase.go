package interfaces

import (
	"context"

	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type CustomerUsecase interface {
	GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error)
	GetCustomerById(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error)
	CreateCustomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error
	UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, id string) error
}
