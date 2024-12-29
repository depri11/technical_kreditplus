package usecase

import (
	"context"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type usecase struct {
	repo   interfaces.CustomerRepository
	tracer trace.Tracer
}

func NewUseCase(repoArticle interfaces.CustomerRepository, cfg *config.Config) *usecase {
	tracer := otel.Tracer(cfg.APP.ServiceName)
	return &usecase{repoArticle, tracer}
}

func (u *usecase) GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error) {
	customer, err := u.repo.GetCustomer(ctx, nik)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (u *usecase) GetCustomerById(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error) {
	customer, err := u.repo.GetCustomer(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (u *usecase) CreateCustomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error {
	return u.repo.CreateCustomer(ctx, customer)
}

func (u *usecase) UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error {
	return u.repo.UpdateCustomer(ctx, customer)
}

func (u *usecase) DeleteCustomer(ctx context.Context, id string) error {
	return u.repo.DeleteCustomer(ctx, id)
}
