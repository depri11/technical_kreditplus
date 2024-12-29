package usecase

import (
	"context"

	transaction_proto "github.com/depri11/technical_kreditplus/protos"
)

type transactionUseCase struct {
	transactionClient transaction_proto.TransactionServiceClient
}

func NewTransactionUseCase(transactionClient transaction_proto.TransactionServiceClient) *transactionUseCase {
	return &transactionUseCase{transactionClient}
}

func (u *transactionUseCase) GetTransaction(ctx context.Context, contractNumber string) (*transaction_proto.GetTransactionResponse, error) {
	result, err := u.transactionClient.GetTransaction(ctx, &transaction_proto.GetTransactionRequest{ContractNumber: contractNumber})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *transactionUseCase) CreateTransaction(ctx context.Context, transaction *transaction_proto.CreateTransactionRequest) error {
	_, err := u.transactionClient.CreateTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (u *transactionUseCase) UpdateTransaction(ctx context.Context, transaction *transaction_proto.UpdateTransactionRequest) error {
	_, err := u.transactionClient.UpdateTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (u *transactionUseCase) DeleteTransaction(ctx context.Context, contractNumner string) error {
	_, err := u.transactionClient.DeleteTransaction(ctx, &transaction_proto.DeleteTransactionRequest{ContractNumber: contractNumner})
	if err != nil {
		return err
	}

	return nil
}
