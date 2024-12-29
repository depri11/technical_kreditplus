package delivery

import (
	"context"

	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"google.golang.org/protobuf/types/known/emptypb"
)

type delivery struct {
	usecase interfaces.CustomerUsecase
}

func NewDelivery(ucArticle interfaces.CustomerUsecase) *delivery {
	return &delivery{ucArticle}
}

func (d *delivery) GetCustomer(ctx context.Context, req *customer_proto.GetCustomerRequest) (*customer_proto.GetCustomerResponse, error) {
	result, err := d.usecase.GetCustomer(ctx, req.Nik)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *delivery) GetCustomerById(ctx context.Context, req *customer_proto.GetCustomerByIdRequest) (*customer_proto.GetCustomerResponse, error) {
	result, err := d.usecase.GetCustomer(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *delivery) CreateCustomer(ctx context.Context, req *customer_proto.CreateCustomerRequest) (*emptypb.Empty, error) {
	err := d.usecase.CreateCustomer(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *delivery) DeleteCustomer(ctx context.Context, req *customer_proto.DeleteCustomerRequest) (*emptypb.Empty, error) {
	err := d.usecase.DeleteCustomer(ctx, req.Nik)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *delivery) UpdateCustomer(ctx context.Context, req *customer_proto.UpdateCustomerRequest) (*emptypb.Empty, error) {
	err := d.usecase.UpdateCustomer(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *delivery) GetCustomerLimit(ctx context.Context, req *customer_proto.GetCustomerLimitRequest) (*customer_proto.GetCustomerLimitResponse, error) {
	result, err := d.usecase.GetCustomerLimit(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *delivery) UpdateCustomerLimit(ctx context.Context, req *customer_proto.UpdateCustomerLimitRequest) (*emptypb.Empty, error) {
	err := d.usecase.UpdateCustomerLimit(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
