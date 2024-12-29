package Repository

import (
	"context"
	"strconv"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/models"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewRepository(db *gorm.DB, cfg *config.Config) *Repository {
	tracer := otel.Tracer(cfg.APP.ServiceName)
	return &Repository{db, tracer}
}

func (r *Repository) withTx(tx *gorm.DB) interfaces.CustomerRepository {
	return &Repository{
		db: tx,
	}
}

func (r *Repository) GetCustomer(ctx context.Context, nik string) (*customer_proto.GetCustomerResponse, error) {
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

func (r *Repository) GetCustomerById(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).Table("customers").Where("id = ?", id).Find(&customer).Error
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

func (r *Repository) CreateCustomer(ctx context.Context, customerProto *customer_proto.CreateCustomerRequest) error {
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

func (r *Repository) UpdateCustomer(ctx context.Context, customerProto *customer_proto.UpdateCustomerRequest) error {
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

func (r *Repository) DeleteCustomer(ctx context.Context, nik string) error {
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

func (r *Repository) GetCustomerLimit(ctx context.Context, id string) (*customer_proto.GetCustomerLimitResponse, error) {
	var customerLimit models.CustomerLimit
	err := r.db.WithContext(ctx).Table("customer_limits").Where("customer_id = ?", id).Find(&customerLimit).Error
	if err != nil {
		return nil, err
	}

	if customerLimit.CustomerId == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	customerProto := &customer_proto.GetCustomerLimitResponse{
		CustomerId:    strconv.Itoa(int(customerLimit.CustomerId)),
		Tenor_1Month:  customerLimit.Tenor1Month,
		Tenor_2Months: customerLimit.Tenor2Month,
		Tenor_3Months: customerLimit.Tenor3Month,
		Tenor_4Months: customerLimit.Tenor4Month,
	}

	return customerProto, nil
}

func (r *Repository) CreateCustomerLimit(ctx context.Context, customerProto *models.CreateCustomerLimit) error {
	customer := &models.CustomerLimit{
		CustomerId:  customerProto.CustomerId,
		Tenor1Month: customerProto.Tenor1Month,
		Tenor2Month: customerProto.Tenor2Month,
		Tenor3Month: customerProto.Tenor3Month,
		Tenor4Month: customerProto.Tenor4Month,
	}
	return r.db.WithContext(ctx).Table("customer_limits").Create(&customer).Error
}

func (r *Repository) UpdateCustomerLimit(ctx context.Context, customerModel *customer_proto.UpdateCustomerLimitRequest) error {
	id, err := strconv.Atoi(customerModel.CustomerId)
	if err != nil {
		return err
	}

	customer := &models.CustomerLimit{
		Tenor1Month: customerModel.Tenor_1Month,
		Tenor2Month: customerModel.Tenor_2Months,
		Tenor3Month: customerModel.Tenor_3Months,
		Tenor4Month: customerModel.Tenor_4Months,
	}
	return r.db.WithContext(ctx).Table("customer_limits").Where("customer_id = ?", id).Updates(customer).Error
}

func (r *Repository) Transaction(ctx context.Context, fn func(repo interfaces.CustomerRepository) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := r.withTx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) DeleteCustomerLimit(ctx context.Context, id string) error {
	customerId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	customer := models.CustomerLimit{
		CustomerId: uint(customerId),
	}
	result := r.db.WithContext(ctx).Table("customer_limits").Where("customer_id = ?", id).Delete(&customer)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
