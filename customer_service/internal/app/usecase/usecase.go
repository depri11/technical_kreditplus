package usecase

import (
	"context"
	"strconv"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
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
	return u.repo.Transaction(ctx, func(repo interfaces.CustomerRepository) error {

		err := repo.CreateCustomer(ctx, customer)
		if err != nil {
			return err
		}

		customer, err := repo.GetCustomer(ctx, customer.Nik)
		if err != nil {
			return err
		}

		id, err := strconv.Atoi(customer.Id)
		if err != nil {
			return err
		}

		repo.CreateCustomerLimit(ctx, &models.CreateCustomerLimit{
			CustomerId: uint(id),
			// TODO: remove this when customer limit is ready
			Tenor1Month: 100000.00,
			Tenor2Month: 200000.00,
			Tenor3Month: 500000.00,
			Tenor4Month: 7000000.00,
		})

		return nil
	})
}

func (u *usecase) UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error {
	return u.repo.UpdateCustomer(ctx, customer)
}

func (u *usecase) DeleteCustomer(ctx context.Context, nik string) error {
	return u.repo.Transaction(ctx, func(repo interfaces.CustomerRepository) error {

		customer, err := repo.GetCustomer(ctx, nik)
		if err != nil {
			return err
		}

		if customer != nil && customer.Id == "" {
			return gorm.ErrRecordNotFound
		}

		err = repo.DeleteCustomerLimit(ctx, customer.Nik)
		if err != nil {
			return err
		}

		err = repo.DeleteCustomer(ctx, nik)
		if err != nil {
			return err
		}
		return nil
	})
}

func (u *usecase) GetCustomerLimit(ctx context.Context, nik string) (*customer_proto.GetCustomerLimitResponse, error) {
	customer, err := u.repo.GetCustomer(ctx, nik)
	if err != nil {
		return nil, err
	}
	customerLimit, err := u.repo.GetCustomerLimit(ctx, customer.Id)
	if err != nil {
		return nil, err
	}
	return customerLimit, nil
}

func (u *usecase) UpdateCustomerLimit(ctx context.Context, customerProto *customer_proto.UpdateCustomerLimitRequest) error {
	return u.repo.UpdateCustomerLimit(ctx, customerProto)
}
