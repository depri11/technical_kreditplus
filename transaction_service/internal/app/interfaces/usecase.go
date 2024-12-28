package interfaces

import (
	"context"

	transaction_proto "github.com/depri11/technical_kreditplus/protos"
)

type TransactionUsecase interface {
	GetTransaction(ctx context.Context, id string) (*transaction_proto.GetTransactionResponse, error)
	CreateTransaction(ctx context.Context, transaction *transaction_proto.CreateTransactionRequest) error
	UpdateTransaction(ctx context.Context, transaction *transaction_proto.UpdateTransactionRequest) error
	DeleteTransaction(ctx context.Context, id string) error
}
