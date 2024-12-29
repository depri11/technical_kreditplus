package interfaces

import (
	"context"

	"github.com/depri11/technical_kreditplus/customer_service/internal/app/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type CustomerRepository interface {
	GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error)
	GetCustomerById(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error)
	CreateCustomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error
	UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error
	DeleteCustomer(ctx context.Context, nik string) error
	GetCustomerLimit(ctx context.Context, id string) (*customer_proto.GetCustomerLimitResponse, error)
	CreateCustomerLimit(ctx context.Context, customerProto *models.CreateCustomerLimit) error
	UpdateCustomerLimit(ctx context.Context, customerModel *customer_proto.UpdateCustomerLimitRequest) error
	DeleteCustomerLimit(ctx context.Context, id string) error

	Transaction(ctx context.Context, fn func(repo CustomerRepository) error) error
}
