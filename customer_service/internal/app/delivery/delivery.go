package delivery

import (
	"context"
	"log"

	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type delivery struct {
	usecase interfaces.CustomerUsecase
}

func NewDelivery(ucArticle interfaces.CustomerUsecase) *delivery {
	return &delivery{ucArticle}
}

func (d *delivery) GetCustomer(ctx context.Context, req *customer_proto.GetCustomerRequest) (*customer_proto.GetCustomerResponse, error) {
	result, err := d.usecase.GetCustomer(ctx, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
