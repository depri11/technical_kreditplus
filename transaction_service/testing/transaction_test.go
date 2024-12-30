package testing_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/depri11/technical_kreditplus/protos"
	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/delivery"
	Repository "github.com/depri11/technical_kreditplus/transaction_service/internal/app/repository"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/usecase"
	customer_mock "github.com/depri11/technical_kreditplus/transaction_service/testing/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func Test_GetTransaction(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "transaction_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "transaction_service",
		},
	}

	repo := Repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, nil, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name       string
		request    *transaction_proto.GetTransactionRequest
		wantResult *transaction_proto.GetTransactionResponse
		wantError  error
		prepare    func()
		cleanUp    func()
	}{
		{
			"Not found Transaction",
			&transaction_proto.GetTransactionRequest{ContractNumber: "123456789"},
			nil,
			gorm.ErrRecordNotFound,
			func() {},
			func() {},
		},
		{
			"Get Transaction Request",
			&transaction_proto.GetTransactionRequest{ContractNumber: "123456789"},
			&transaction_proto.GetTransactionResponse{
				CustomerId:     "CUSTOMERID-1",
				ContractNumber: "123456789",
			},
			nil,
			func() {
				err := repo.CreateTransaction(context.Background(), &transaction_proto.CreateTransactionRequest{
					CustomerId:        "CUSTOMERID-1",
					ContractNumber:    "123456789",
					Otr:               5000000.00,
					AdminFee:          500000.00,
					InstallmentAmount: 500000.00,
					InterestAmount:    500000.00,
					AssetName:         "Asset 1",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteTransaction(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.GetTransaction(context.Background(), &transaction_proto.GetTransactionRequest{ContractNumber: "123456789"})
			if res != nil {
				assert.Equal(t, res.CustomerId, tt.wantResult.CustomerId)
				assert.Equal(t, res.ContractNumber, tt.wantResult.ContractNumber)
			}
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func Test_DeleteTransaction(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "transaction_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "transaction_service",
		},
	}

	repo := Repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, nil, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name       string
		request    *transaction_proto.DeleteTransactionRequest
		wantResult *emptypb.Empty
		wantError  error
		prepare    func()
		cleanUp    func()
	}{
		{
			"Not found Transaction",
			&transaction_proto.DeleteTransactionRequest{ContractNumber: "123456789"},
			nil,
			gorm.ErrRecordNotFound,
			func() {
				repo.DeleteTransaction(context.Background(), "123456789")
			},
			func() {},
		},
		{
			"Delete Transaction Request",
			&transaction_proto.DeleteTransactionRequest{ContractNumber: "123456789"},
			&emptypb.Empty{},
			nil,
			func() {
				err := repo.CreateTransaction(context.Background(), &transaction_proto.CreateTransactionRequest{
					CustomerId:        "CUSTOMERID-1",
					ContractNumber:    "123456789",
					Otr:               5000000.00,
					AdminFee:          500000.00,
					InstallmentAmount: 500000.00,
					InterestAmount:    500000.00,
					AssetName:         "Asset 1",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteTransaction(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.DeleteTransaction(context.Background(), &transaction_proto.DeleteTransactionRequest{ContractNumber: "123456789"})
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantError, err)

			_, err = delivery.GetTransaction(context.Background(), &transaction_proto.GetTransactionRequest{ContractNumber: "123456789"})
			assert.Equal(t, gorm.ErrRecordNotFound, err)
		})
	}
}

func Test_UpdateTransaction(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "transaction_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "transaction_service",
		},
	}

	repo := Repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, nil, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name            string
		request         *transaction_proto.UpdateTransactionRequest
		wantResult      *emptypb.Empty
		wantError       error
		WantTransaction *transaction_proto.GetTransactionResponse
		prepare         func()
		cleanUp         func()
	}{
		{
			"Not found Transaction",
			&transaction_proto.UpdateTransactionRequest{
				Id:                "1",
				CustomerId:        "CUSTOMERID-1",
				ContractNumber:    "123456789",
				Otr:               5000000.00,
				AdminFee:          500000.00,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "Asset 1",
			},
			nil,
			gorm.ErrRecordNotFound,
			nil,
			func() {},
			func() {},
		},
		{
			"Update Transaction Request",
			&transaction_proto.UpdateTransactionRequest{
				CustomerId:        "CUSTOMERID-1",
				ContractNumber:    "123456789",
				Otr:               1,
				AdminFee:          1,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "EDIT Asset 1",
			},
			&emptypb.Empty{},
			nil,
			&transaction_proto.GetTransactionResponse{
				CustomerId:        "CUSTOMERID-1",
				ContractNumber:    "123456789",
				Otr:               1,
				AdminFee:          1,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "EDIT Asset 1",
			},
			func() {
				err := repo.CreateTransaction(context.Background(), &transaction_proto.CreateTransactionRequest{
					CustomerId:        "CUSTOMERID-1",
					ContractNumber:    "123456789",
					Otr:               5000000.00,
					AdminFee:          500000.00,
					InstallmentAmount: 500000.00,
					InterestAmount:    500000.00,
					AssetName:         "Asset 1",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteTransaction(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.UpdateTransaction(context.Background(), tt.request)
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantError, err)

			transaction, err := delivery.GetTransaction(context.Background(), &transaction_proto.GetTransactionRequest{ContractNumber: "123456789"})
			assert.Equal(t, tt.wantError, err)
			if transaction != nil {
				assert.Equal(t, tt.WantTransaction.Otr, transaction.Otr)
				assert.Equal(t, tt.WantTransaction.AdminFee, transaction.AdminFee)
				assert.Equal(t, tt.WantTransaction.AssetName, transaction.AssetName)
			}
		})
	}
}

func Test_CreateTransaction(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "transaction_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "transaction_service",
		},
	}

	// Create a new instance of the mock customer service
	mockCustomerService := new(customer_mock.MockCustomerService)

	repo := Repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, mockCustomerService, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name      string
		request   *transaction_proto.CreateTransactionRequest
		wantError error
		prepare   func()
		cleanUp   func()
	}{
		{
			"Invalid tenor",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber: "123456789",
			},
			errors.New("Invalid tenor"),
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{
					Id:           "1",
					FullName:     "Devri",
					Nik:          "123456789",
					LegalName:    "Devri",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Now()),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(&protos.GetCustomerLimitResponse{
					CustomerId:   "1",
					Tenor_1Month: 10000000000.00,
				}, nil)
				mockCustomerService.On("UpdateCustomerLimit", mock.Anything, &protos.UpdateCustomerLimitRequest{Tenor_1Month: 9998000000}).Return(nil, nil)
			},
			func() {
			},
		},
		{
			"Customer And Customer Limit Not Found",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber: "123456789",
			},
			errors.New("Customer not found"),
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(&protos.GetCustomerLimitResponse{}, nil)
			},
			func() {
			},
		},
		{
			"Customer Limit Not Found",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber: "123456789",
			},
			errors.New("Customer limit not found"),
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{
					Id:           "1",
					FullName:     "Devri",
					Nik:          "123456789",
					LegalName:    "Devri",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Now()),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(nil, nil)
			},
			func() {
			},
		},
		{
			"Tenor limit does not meet the minimum",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber:    "123456789",
				Otr:               500000.00,
				AdminFee:          500000.00,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "Rumah",
				Tenor:             protos.Tenor_TENOR_1_MONTH,
			},
			errors.New("Transaction amount exceeds the limit for 1 month tenor"),
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{
					Id:           "1",
					FullName:     "Devri",
					Nik:          "123456789",
					LegalName:    "Devri",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Now()),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(&protos.GetCustomerLimitResponse{
					CustomerId:   "1",
					Tenor_1Month: 1.00,
				}, nil)
			},
			func() {
			},
		},
		{
			"Failed update customer limit",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber:    "123456789",
				Otr:               500000.00,
				AdminFee:          500000.00,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "Rumah",
				Tenor:             protos.Tenor_TENOR_1_MONTH,
			},
			errors.New("error updating customer limit"),
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{
					Id:           "1",
					FullName:     "Devri",
					Nik:          "123456789",
					LegalName:    "Devri",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Now()),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(&protos.GetCustomerLimitResponse{
					CustomerId:   "1",
					Tenor_1Month: 10000000000.00,
				}, nil)
				mockCustomerService.On("UpdateCustomerLimit", mock.Anything, &protos.UpdateCustomerLimitRequest{Tenor_1Month: 9998000000}).Return(nil, errors.New("error updating customer limit"))
			},
			func() {
			},
		},
		{
			"Create Transaction Request",
			&transaction_proto.CreateTransactionRequest{
				ContractNumber:    "123456789",
				Otr:               500000.00,
				AdminFee:          500000.00,
				InstallmentAmount: 500000.00,
				InterestAmount:    500000.00,
				AssetName:         "Rumah",
				Tenor:             protos.Tenor_TENOR_1_MONTH,
			},
			nil,
			func() {
				mockCustomerService.On("GetCustomerById", mock.Anything, &protos.GetCustomerByIdRequest{Id: ""}).Return(&protos.GetCustomerResponse{
					Id:           "1",
					FullName:     "Devri",
					Nik:          "123456789",
					LegalName:    "Devri",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Now()),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				}, nil)
				mockCustomerService.On("GetCustomerLimit", mock.Anything, &protos.GetCustomerLimitRequest{CustomerId: ""}).Return(&protos.GetCustomerLimitResponse{
					CustomerId:   "1",
					Tenor_1Month: 10000000000.00,
				}, nil)
				mockCustomerService.On("UpdateCustomerLimit", mock.Anything, &protos.UpdateCustomerLimitRequest{Tenor_1Month: 9998000000}).Return(nil, nil)
			},
			func() {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			_, err := delivery.CreateTransaction(context.Background(), tt.request)
			assert.Equal(t, tt.wantError, err)

			mockCustomerService.ExpectedCalls = nil
			mockCustomerService.Calls = nil

			tt.cleanUp()
		})
	}
}
