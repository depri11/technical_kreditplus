package delivery

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/models"
	"github.com/depri11/technical_kreditplus/api_gateway/internal/transaction/interfaces"
	trx_model "github.com/depri11/technical_kreditplus/api_gateway/internal/transaction/models"
	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/gorilla/mux"
)

type transactionDelivery struct {
	r                  *mux.Router
	transactionUseCase interfaces.TransactionUsecase
}

func NewDelivery(transactionUseCase interfaces.TransactionUsecase, r *mux.Router) *transactionDelivery {
	return &transactionDelivery{transactionUseCase: transactionUseCase, r: r}
}

func (d *transactionDelivery) GetTransaction(w http.ResponseWriter, r *http.Request) {
	contractNumber := mux.Vars(r)["contract_number"]

	ctx := context.Background()

	result := models.GeneralResponse[*trx_model.GetTransactionResponse]{
		Success: false,
	}

	transaction, err := d.transactionUseCase.GetTransaction(ctx, contractNumber)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	data := &trx_model.GetTransactionResponse{
		Id:                transaction.Id,
		CustomerId:        transaction.CustomerId,
		ContractNumber:    transaction.ContractNumber,
		Otr:               transaction.Otr,
		AdminFee:          transaction.AdminFee,
		InstallmentAmount: transaction.InstallmentAmount,
		InterestAmount:    transaction.InterestAmount,
		AssetName:         transaction.AssetName,
		CreatedAt:         transaction.CreatedAt.AsTime(),
		UpdatedAt:         transaction.UpdatedAt.AsTime(),
	}

	result.Message = "Success"
	result.Success = true
	result.Data = data
	result.ToJson(w)
}

func (d *transactionDelivery) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var payload transaction_proto.CreateTransactionRequest

	result := models.GeneralResponse[*transaction_proto.GetTransactionResponse]{
		Success: false,
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	ctx := r.Context()

	err = d.transactionUseCase.CreateTransaction(ctx, &payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Create Transaction"
	result.Success = true
	result.ToJson(w)
}

func (d *transactionDelivery) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var payload transaction_proto.UpdateTransactionRequest

	result := models.GeneralResponse[*transaction_proto.UpdateTransactionRequest]{
		Success: false,
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	ctx := r.Context()

	err = d.transactionUseCase.UpdateTransaction(ctx, &payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Update Transaction"
	result.Success = true
	result.Data = &payload
	result.ToJson(w)
}

func (d *transactionDelivery) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	contractNumber := mux.Vars(r)["contract_number"]

	ctx := r.Context()

	result := models.GeneralResponse[*transaction_proto.UpdateTransactionRequest]{
		Success: false,
	}

	err := d.transactionUseCase.DeleteTransaction(ctx, contractNumber)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Delete Transaction"
	result.Success = true
	result.ToJson(w)
}
