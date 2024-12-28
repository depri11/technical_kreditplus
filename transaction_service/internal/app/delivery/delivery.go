package delivery

import (
	"context"

	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/interfaces"
	"google.golang.org/protobuf/types/known/emptypb"
)

type delivery struct {
	usecase interfaces.TransactionUsecase
}

func NewDelivery(ucArticle interfaces.TransactionUsecase) *delivery {
	return &delivery{ucArticle}
}

func (d *delivery) GetTransaction(ctx context.Context, req *transaction_proto.GetTransactionRequest) (*transaction_proto.GetTransactionResponse, error) {
	result, err := d.usecase.GetTransaction(ctx, req.ContractNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *delivery) CreateTransaction(ctx context.Context, req *transaction_proto.CreateTransactionRequest) (*emptypb.Empty, error) {
	err := d.usecase.CreateTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *delivery) DeleteTransaction(ctx context.Context, req *transaction_proto.DeleteTransactionRequest) (*emptypb.Empty, error) {
	err := d.usecase.DeleteTransaction(ctx, req.ContractNumber)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *delivery) UpdateTransaction(ctx context.Context, req *transaction_proto.UpdateTransactionRequest) (*emptypb.Empty, error) {
	err := d.usecase.UpdateTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
