package interfaces

import (
	"context"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/customer/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type CustomerUsecase interface {
	GetCustomer(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error)
	CreateCustomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error
	UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, nik string) error

	GetCustomerLimit(ctx context.Context, id string) (*customer_proto.GetCustomerLimitResponse, error)
	UpdateCustomerLimit(ctx context.Context, req *models.UpdateCustomerLimitRequest) error
}
