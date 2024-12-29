package usecase

import (
	"context"

	"github.com/depri11/technical_kreditplus/api_gateway/internal/customer/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type customerUseCase struct {
	customerClient customer_proto.CustomerServiceClient
}

func NewCustomerUseCase(customerClient customer_proto.CustomerServiceClient) *customerUseCase {
	return &customerUseCase{customerClient}
}

func (u *customerUseCase) GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error) {
	result, err := u.customerClient.GetCustomer(ctx, &customer_proto.GetCustomerRequest{Nik: nik})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUseCase) CreateCustomer(ctx context.Context, customer *customer_proto.CreateCustomerRequest) error {
	_, err := u.customerClient.CreateCustomer(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (u *customerUseCase) UpdateCustomer(ctx context.Context, customer *customer_proto.UpdateCustomerRequest) error {
	_, err := u.customerClient.UpdateCustomer(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (u *customerUseCase) DeleteCustomer(ctx context.Context, nik string) error {
	_, err := u.customerClient.DeleteCustomer(ctx, &customer_proto.DeleteCustomerRequest{Nik: nik})
	if err != nil {
		return err
	}

	return nil
}

func (u *customerUseCase) GetCustomerLimit(ctx context.Context, id string) (*customer_proto.GetCustomerLimitResponse, error) {
	result, err := u.customerClient.GetCustomerLimit(ctx, &customer_proto.GetCustomerLimitRequest{CustomerId: id})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *customerUseCase) UpdateCustomerLimit(ctx context.Context, req *models.UpdateCustomerLimitRequest) error {

	customer, err := u.customerClient.GetCustomer(ctx, &customer_proto.GetCustomerRequest{Nik: req.Nik})
	if err != nil {
		return err
	}

	customerLimit := &customer_proto.UpdateCustomerLimitRequest{
		CustomerId:    customer.Id,
		Tenor_1Month:  req.Tenor1Month,
		Tenor_2Months: req.Tenor2Month,
		Tenor_3Months: req.Tenor3Month,
		Tenor_4Months: req.Tenor4Month,
	}
	_, err = u.customerClient.UpdateCustomerLimit(ctx, customerLimit)
	if err != nil {
		return err
	}

	return nil
}
