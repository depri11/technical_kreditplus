package testing_test

import (
	"context"
	"testing"
	"time"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/delivery"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/repository"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/usecase"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func Test_GetCustomer(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "customer_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "customer_service",
		},
	}

	repo := repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name       string
		request    *customer_proto.GetCustomerRequest
		wantResult *customer_proto.GetCustomerResponse
		wantError  error
		prepare    func()
		cleanUp    func()
	}{
		{
			"Not found Customer",
			&customer_proto.GetCustomerRequest{Nik: "123456789"},
			nil,
			gorm.ErrRecordNotFound,
			func() {},
			func() {},
		},
		{
			"Get Customer Request",
			&customer_proto.GetCustomerRequest{Nik: "123456789"},
			&customer_proto.GetCustomerResponse{
				Id:           "1",
				Nik:          "123456789",
				FullName:     "CUSTOMER 1",
				LegalName:    "CUSTOMER 1",
				PlaceOfBirth: "Jakarta",
				DateOfBirth:  timestamppb.New(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Salary:       5000000.00,
				PhotoKtp:     "ktp.jpg",
				PhotoSelfie:  "selfie.jpg",
			},
			nil,
			func() {
				dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
				// Convert time.Time to *timestamp.Timestamp
				dobTimestamp := timestamppb.New(dob)

				err := repo.CreateCustomer(context.Background(), &customer_proto.CreateCustomerRequest{
					Nik:          "123456789",
					FullName:     "CUSTOMER 1",
					LegalName:    "CUSTOMER 1",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  dobTimestamp,
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteCustomer(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.GetCustomer(context.Background(), &customer_proto.GetCustomerRequest{Nik: "123456789"})
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantError, err)
		})
	}
}

func Test_DeleteCustomer(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "customer_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "customer_service",
		},
	}

	repo := repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name       string
		request    *customer_proto.DeleteCustomerRequest
		wantResult *emptypb.Empty
		wantError  error
		prepare    func()
		cleanUp    func()
	}{
		{
			"Not found Customer",
			&customer_proto.DeleteCustomerRequest{Nik: "123456789"},
			nil,
			gorm.ErrRecordNotFound,
			func() {
				repo.DeleteCustomer(context.Background(), "123456789")
			},
			func() {},
		},
		{
			"Delete Customer Request",
			&customer_proto.DeleteCustomerRequest{Nik: "123456789"},
			&emptypb.Empty{},
			nil,
			func() {
				dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
				// Convert time.Time to *timestamp.Timestamp
				dobTimestamp := timestamppb.New(dob)

				err := repo.CreateCustomer(context.Background(), &customer_proto.CreateCustomerRequest{
					Nik:          "123456789",
					FullName:     "CUSTOMER 1",
					LegalName:    "CUSTOMER 1",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  dobTimestamp,
					Salary:       50000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteCustomer(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.DeleteCustomer(context.Background(), &customer_proto.DeleteCustomerRequest{Nik: "123456789"})
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantError, err)

			_, err = delivery.GetCustomer(context.Background(), &customer_proto.GetCustomerRequest{Nik: "123456789"})
			assert.Equal(t, gorm.ErrRecordNotFound, err)
		})
	}
}

func Test_UpdateCustomer(t *testing.T) {
	cfg := &config.Config{
		APP: config.App{
			ServiceName: "customer_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "passwor",
			Host:     "localhost",
			Port:     5543,
			Name:     "customer_service",
		},
	}

	repo := repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, cfg)
	delivery := delivery.NewDelivery(usecase)

	tests := []struct {
		name         string
		request      *customer_proto.UpdateCustomerRequest
		wantResult   *emptypb.Empty
		wantError    error
		WantCustomer *customer_proto.GetCustomerResponse
		prepare      func()
		cleanUp      func()
	}{
		{
			"Not found Customer",
			&customer_proto.UpdateCustomerRequest{
				Id:           "1",
				Nik:          "123456789",
				FullName:     "EDIT CUSTOMER 1",
				LegalName:    "EDIT CUSTOMER 1",
				PlaceOfBirth: "Jakarta",
				DateOfBirth:  nil,
				Salary:       0.0,
				PhotoKtp:     "ktp.jpg",
				PhotoSelfie:  "selfie.jpg",
			},
			nil,
			gorm.ErrRecordNotFound,
			nil,
			func() {},
			func() {},
		},
		{
			"Update Customer Request",
			&customer_proto.UpdateCustomerRequest{
				Id:           "1",
				Nik:          "123456789",
				FullName:     "EDIT FULLNAME 1",
				LegalName:    "EDIT FULLNAME 1",
				PlaceOfBirth: "Jakarta",
				DateOfBirth:  timestamppb.New(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Salary:       5000000.00,
				PhotoKtp:     "ktp.jpg",
				PhotoSelfie:  "selfie.jpg",
			},
			&emptypb.Empty{},
			nil,
			&customer_proto.GetCustomerResponse{
				Nik:          "123456789",
				FullName:     "EDIT FULLNAME 1",
				LegalName:    "EDIT FULLNAME 1",
				PlaceOfBirth: "Jakarta",
				DateOfBirth:  timestamppb.New(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Salary:       5000000.00,
				PhotoKtp:     "ktp.jpg",
				PhotoSelfie:  "selfie.jpg",
			},
			func() {
				err := repo.CreateCustomer(context.Background(), &customer_proto.CreateCustomerRequest{
					Nik:          "123456789",
					FullName:     "CUSTOMER 1",
					LegalName:    "CUSTOMER 1",
					PlaceOfBirth: "Jakarta",
					DateOfBirth:  timestamppb.New(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)),
					Salary:       5000000.00,
					PhotoKtp:     "ktp.jpg",
					PhotoSelfie:  "selfie.jpg",
				})
				assert.Nil(t, err)
			},
			func() {
				repo.DeleteCustomer(context.Background(), "123456789")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			res, err := delivery.UpdateCustomer(context.Background(), tt.request)
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantError, err)

			customer, err := delivery.GetCustomer(context.Background(), &customer_proto.GetCustomerRequest{Nik: "123456789"})
			assert.Equal(t, tt.wantError, err)
			if customer != nil {
				assert.Equal(t, tt.WantCustomer.Nik, customer.Nik)
				assert.Equal(t, tt.WantCustomer.FullName, customer.FullName)
				assert.Equal(t, tt.WantCustomer.LegalName, customer.LegalName)
			}
		})
	}
}
