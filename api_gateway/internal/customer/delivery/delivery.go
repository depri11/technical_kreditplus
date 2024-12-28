package delivery

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/customer/interfaces"
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
	id := mux.Vars(r)["id"]

	ctx := context.Background()
	result, err := d.customerUseCase.GetCustomer(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resByte, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resByte)
}

func (d *customerDelivery) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var payload customer_proto.CreateCustomerRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	err = d.customerUseCase.CreateCustomer(ctx, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}
