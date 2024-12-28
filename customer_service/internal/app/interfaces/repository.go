package interfaces

import (
	"context"

	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type CustomerRepository interface {
	GetCustomer(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error)
}
