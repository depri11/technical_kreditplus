package usecase

import (
	"context"
	"errors"

	"github.com/depri11/technical_kreditplus/protos"
	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/interfaces"
	"github.com/shopspring/decimal"
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
		Id: transaction.CustomerId,
	}
	customer, err := u.customerRepo.GetCustomerById(ctx, req)
	if err != nil {
		return err
	}

	if customer == nil || customer.Nik == "" {
		return errors.New("Customer not found")
	}

	// check customer limits
	limitReq := &transaction_proto.GetCustomerLimitRequest{
		CustomerId: transaction.CustomerId,
	}
	customerLimit, err := u.customerRepo.GetCustomerLimit(ctx, limitReq)
	if err != nil {
		return err
	}

	if customerLimit == nil {
		return errors.New("Customer limit not found")
	}

	// Calculate the total transaction amount
	totalTransactionAmount := decimal.NewFromFloat(transaction.Otr).Add(decimal.NewFromFloat(transaction.AdminFee)).Add(decimal.NewFromFloat(transaction.InstallmentAmount)).Add(decimal.NewFromFloat(transaction.InterestAmount))

	// Check if the transaction amount is within the customer's limit for the given tenor
	switch transaction.Tenor {
	case protos.Tenor_TENOR_1_MONTH:
		if totalTransactionAmount.GreaterThan(decimal.NewFromFloat(customerLimit.Tenor_1Month)) {
			return errors.New("Transaction amount exceeds the limit for 1 month tenor")
		}

		newLimit := decimal.NewFromFloat(customerLimit.Tenor_1Month).Sub(totalTransactionAmount)
		customerLimit.Tenor_1Month = newLimit.InexactFloat64()
	case protos.Tenor_TENOR_2_MONTH:
		if totalTransactionAmount.GreaterThan(decimal.NewFromFloat(customerLimit.Tenor_2Months)) {
			return errors.New("Transaction amount exceeds the limit for 2 months tenor")
		}

		newLimit := decimal.NewFromFloat(customerLimit.Tenor_2Months).Sub(totalTransactionAmount)
		customerLimit.Tenor_2Months = newLimit.InexactFloat64()
	case protos.Tenor_TENOR_3_MONTH:
		if totalTransactionAmount.GreaterThan(decimal.NewFromFloat(customerLimit.Tenor_3Months)) {
			return errors.New("Transaction amount exceeds the limit for 3 months tenor")
		}

		newLimit := decimal.NewFromFloat(customerLimit.Tenor_3Months).Sub(totalTransactionAmount)
		customerLimit.Tenor_3Months = newLimit.InexactFloat64()
	case protos.Tenor_TENOR_4_MONTH:
		if totalTransactionAmount.GreaterThan(decimal.NewFromFloat(customerLimit.Tenor_4Months)) {
			return errors.New("Transaction amount exceeds the limit for 4 months tenor")
		}

		newLimit := decimal.NewFromFloat(customerLimit.Tenor_4Months).Sub(totalTransactionAmount)
		customerLimit.Tenor_4Months = newLimit.InexactFloat64()
	default:
		return errors.New("Invalid tenor")
	}

	newCustomerLimit := &protos.UpdateCustomerLimitRequest{
		CustomerId:    transaction.CustomerId,
		Tenor_1Month:  customerLimit.Tenor_1Month,
		Tenor_2Months: customerLimit.Tenor_2Months,
		Tenor_3Months: customerLimit.Tenor_3Months,
		Tenor_4Months: customerLimit.Tenor_4Months,
	}

	// Create the transaction
	return u.repo.Transaction(ctx, func(repo interfaces.TransactionRepository) error {
		err := repo.CreateTransaction(ctx, transaction)
		if err != nil {
			return err
		}

		// Update the customer limit in the database
		_, err = u.customerRepo.UpdateCustomerLimit(ctx, newCustomerLimit)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *usecase) UpdateTransaction(ctx context.Context, transaction *transaction_proto.UpdateTransactionRequest) error {
	return u.repo.UpdateTransaction(ctx, transaction)
}

func (u *usecase) DeleteTransaction(ctx context.Context, id string) error {
	return u.repo.DeleteTransaction(ctx, id)
}
