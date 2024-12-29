package delivery

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/models"
	"github.com/depri11/technical_kreditplus/api_gateway/internal/transaction/interfaces"
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
	nik := mux.Vars(r)["nik"]

	ctx := context.Background()

	result := models.GeneralResponse[*transaction_proto.GetTransactionResponse]{
		Success: false,
	}

	transaction, err := d.transactionUseCase.GetTransaction(ctx, nik)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success"
	result.Success = true
	result.Data = transaction
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
	nik := mux.Vars(r)["nik"]

	ctx := r.Context()

	result := models.GeneralResponse[*transaction_proto.UpdateTransactionRequest]{
		Success: false,
	}

	err := d.transactionUseCase.DeleteTransaction(ctx, nik)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Delete Transaction"
	result.Success = true
	result.ToJson(w)
}
