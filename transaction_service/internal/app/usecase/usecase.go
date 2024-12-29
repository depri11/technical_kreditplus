package usecase

import (
	"context"
	"errors"

	"github.com/depri11/technical_kreditplus/protos"
	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/interfaces"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type usecase struct {
	repo         interfaces.TransactionRepository
	customerRepo protos.CustomerServiceClient
	tracer       trace.Tracer
}

func NewUseCase(repo interfaces.TransactionRepository, customerRepo protos.CustomerServiceClient, cfg *config.Config) *usecase {
	tracer := otel.Tracer(cfg.APP.ServiceName)
	return &usecase{repo, customerRepo, tracer}
}

func (u *usecase) GetTransaction(ctx context.Context, id string) (*transaction_proto.GetTransactionResponse, error) {
	transaction, err := u.repo.GetTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (u *usecase) CreateTransaction(ctx context.Context, transaction *transaction_proto.CreateTransactionRequest) error {
	// check customer data
	req := &transaction_proto.GetCustomerByIdRequest{
		Id: transaction.Id,
	}
	customer, err := u.customerRepo.GetCustomerById(ctx, req)
	if err != nil {
		return err
	}

	if customer == nil || customer.Nik == "" {
		return errors.New("Customer not found")
	}

	return u.repo.CreateTransaction(ctx, transaction)
}

func (u *usecase) UpdateTransaction(ctx context.Context, transaction *transaction_proto.UpdateTransactionRequest) error {
	return u.repo.UpdateTransaction(ctx, transaction)
}

func (u *usecase) DeleteTransaction(ctx context.Context, id string) error {
	return u.repo.DeleteTransaction(ctx, id)
}
