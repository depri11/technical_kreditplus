package repository

import (
	"context"
	"strconv"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewRepository(db *gorm.DB, cfg *config.Config) *repository {
	tracer := otel.Tracer(cfg.APP.ServiceName)
	return &repository{db, tracer}
}

func (r *repository) GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).Table("customers").Where("nik = ?", nik).Find(&customer).Error
	if err != nil {
		return nil, err
	}

	if customer.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	customerProto := &customer_proto.GetCustomerResponse{
		Id:           strconv.Itoa(int(customer.Id)),
		Nik:          customer.Nik,
		FullName:     customer.FullName,
		LegalName:    customer.LegalName,
		PlaceOfBirth: customer.PlaceOfBirth,
		DateOfBirth:  timestamppb.New(customer.DateOfBirth),
		Salary:       customer.Salary,
		PhotoKtp:     customer.PhotoKtp,
		PhotoSelfie:  customer.PhotoSelfie,
	}

	return customerProto, nil
}

func (r *repository) CreateCustomer(ctx context.Context, customerProto *customer_proto.CreateCustomerRequest) error {
	customer := &models.Customer{
		Nik:          customerProto.Nik,
		FullName:     customerProto.FullName,
		LegalName:    customerProto.LegalName,
		PlaceOfBirth: customerProto.PlaceOfBirth,
		DateOfBirth:  customerProto.DateOfBirth.AsTime(),
		Salary:       customerProto.Salary,
		PhotoKtp:     customerProto.PhotoKtp,
		PhotoSelfie:  customerProto.PhotoSelfie,
	}
	return r.db.WithContext(ctx).Table("customers").Create(&customer).Error
}

func (r *repository) UpdateCustomer(ctx context.Context, customerProto *customer_proto.UpdateCustomerRequest) error {
	customerDB, err := r.GetCustomer(ctx, customerProto.Nik)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(customerDB.Id)
	if err != nil {
		return err
	}

	customer := models.Customer{
		Id:           uint(id),
		Nik:          customerProto.Nik,
		FullName:     customerProto.FullName,
		LegalName:    customerProto.LegalName,
		PlaceOfBirth: customerProto.PlaceOfBirth,
		DateOfBirth:  customerProto.DateOfBirth.AsTime(),
		Salary:       customerProto.Salary,
		PhotoKtp:     customerProto.PhotoKtp,
		PhotoSelfie:  customerProto.PhotoSelfie,
	}

	result := r.db.WithContext(ctx).Table("customers").Save(&customer)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) DeleteCustomer(ctx context.Context, nik string) error {
	customerDB, err := r.GetCustomer(ctx, nik)
	if err != nil {
		return err
	}

	customerId, err := strconv.Atoi(customerDB.Id)
	if err != nil {
		return err
	}

	customer := models.Customer{
		Id:           uint(customerId),
		Nik:          customerDB.Nik,
		FullName:     customerDB.FullName,
		LegalName:    customerDB.LegalName,
		PlaceOfBirth: customerDB.PlaceOfBirth,
		DateOfBirth:  customerDB.DateOfBirth.AsTime(),
		Salary:       customerDB.Salary,
		PhotoKtp:     customerDB.PhotoKtp,
		PhotoSelfie:  customerDB.PhotoSelfie,
	}

	result := r.db.WithContext(ctx).Table("customers").Where("nik = ?", nik).Delete(&customer)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
