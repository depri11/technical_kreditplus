package usecase

import (
	"context"

	customer_proto "github.com/depri11/technical_kreditplus/protos"
)

type customerUseCase struct {
	customerClient customer_proto.CustomerServiceClient
}

func NewArticleUseCase(customerClient customer_proto.CustomerServiceClient) *customerUseCase {
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
