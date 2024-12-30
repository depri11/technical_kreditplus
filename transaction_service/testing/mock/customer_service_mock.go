package mock

import (
	"context"

	"github.com/depri11/technical_kreditplus/protos"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) GetCustomerById(ctx context.Context, req *protos.GetCustomerByIdRequest, opts ...grpc.CallOption) (*protos.GetCustomerResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*protos.GetCustomerResponse), args.Error(1)
}

func (m *MockCustomerService) GetCustomerLimit(ctx context.Context, req *protos.GetCustomerLimitRequest, opts ...grpc.CallOption) (*protos.GetCustomerLimitResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*protos.GetCustomerLimitResponse), args.Error(1)
}

func (m *MockCustomerService) UpdateCustomerLimit(ctx context.Context, req *protos.UpdateCustomerLimitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockCustomerService) GetCustomer(ctx context.Context, req *protos.GetCustomerRequest, opts ...grpc.CallOption) (*protos.GetCustomerResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*protos.GetCustomerResponse), args.Error(1)
}

func (m *MockCustomerService) CreateCustomer(ctx context.Context, req *protos.CreateCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockCustomerService) UpdateCustomer(ctx context.Context, req *protos.UpdateCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockCustomerService) DeleteCustomer(ctx context.Context, req *protos.DeleteCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}
