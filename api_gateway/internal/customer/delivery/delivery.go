package delivery

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/customer/interfaces"
	customerModels "github.com/depri11/technical_kreditplus/api_gateway/internal/customer/models"
	"github.com/depri11/technical_kreditplus/api_gateway/internal/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/gorilla/mux"
)

type customerDelivery struct {
	r               *mux.Router
	customerUseCase interfaces.CustomerUsecase
}

func NewDelivery(customerUseCase interfaces.CustomerUsecase, r *mux.Router) *customerDelivery {
	return &customerDelivery{customerUseCase: customerUseCase, r: r}
}

func (d *customerDelivery) GetCustomer(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]

	ctx := context.Background()

	result := models.GeneralResponse[*customerModels.GetCustomerResponse]{
		Success: false,
	}

	customer, err := d.customerUseCase.GetCustomer(ctx, nik)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success"
	result.Success = true
	result.Data = &customerModels.GetCustomerResponse{
		Nik:          customer.Nik,
		FullName:     customer.FullName,
		LegalName:    customer.LegalName,
		PlaceOfBirth: customer.PlaceOfBirth,
		DateOfBirth:  customer.DateOfBirth.AsTime().Format("2006-01-02"),
		Salary:       customer.Salary,
		PhotoKtp:     customer.PhotoKtp,
		PhotoSelfie:  customer.PhotoSelfie,
	}
	result.ToJson(w)
}

func (d *customerDelivery) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var payload customerModels.CreateCustomerRequest

	result := models.GeneralResponse[*customer_proto.GetCustomerResponse]{
		Success: false,
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	protoModel, err := payload.ToProto()
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	ctx := r.Context()

	err = d.customerUseCase.CreateCustomer(ctx, protoModel)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Create Customer"
	result.Success = true
	result.ToJson(w)
}

func (d *customerDelivery) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]

	var payload customerModels.UpdateCustomerRequest

	result := models.GeneralResponse[*customerModels.UpdateCustomerRequest]{
		Success: false,
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	payload.Nik = nik

	ctx := r.Context()
	transactionProto, err := payload.ToProto()
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}
	err = d.customerUseCase.UpdateCustomer(ctx, transactionProto)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Update Customer"
	result.Success = true
	result.Data = &payload
	result.ToJson(w)
}

func (d *customerDelivery) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	nik := mux.Vars(r)["nik"]

	ctx := r.Context()

	result := models.GeneralResponse[*customer_proto.UpdateCustomerRequest]{
		Success: false,
	}

	err := d.customerUseCase.DeleteCustomer(ctx, nik)
	if err != nil {
		result.Message = err.Error()
		result.ToJson(w)
		return
	}

	result.Message = "Success Delete Customer"
	result.Success = true
	result.ToJson(w)
}
