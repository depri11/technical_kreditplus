package Repository

import (
	"context"
	"strconv"

	transaction_proto "github.com/depri11/technical_kreditplus/protos"
	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/interfaces"
	"github.com/depri11/technical_kreditplus/transaction_service/internal/app/models"
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

func (r *Repository) withTx(tx *gorm.DB) interfaces.TransactionRepository {
	return &Repository{
		db: tx,
	}
}

func (r *Repository) GetTransaction(ctx context.Context, contractNumber string) (*transaction_proto.GetTransactionResponse, error) {
	var transaction models.Transaction
	err := r.db.WithContext(ctx).Table("transactions").Where("contract_number = ?", contractNumber).Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	if transaction.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	transactionProto := &transaction_proto.GetTransactionResponse{
		Id:                strconv.Itoa(int(transaction.Id)),
		CustomerId:        transaction.CustomerId,
		ContractNumber:    transaction.ContractNumber,
		Otr:               transaction.Otr,
		AdminFee:          transaction.AdminFee,
		InstallmentAmount: transaction.InstallmentAmount,
		InterestAmount:    transaction.InterestAmount,
		AssetName:         transaction.AssetName,
		CreatedAt:         timestamppb.New(transaction.CreatedAt),
		UpdatedAt:         timestamppb.New(transaction.UpdatedAt),
	}

	return transactionProto, nil
}

func (r *Repository) CreateTransaction(ctx context.Context, transactionProto *transaction_proto.CreateTransactionRequest) error {
	transaction := &models.Transaction{
		ContractNumber:    transactionProto.ContractNumber,
		CustomerId:        transactionProto.CustomerId,
		Otr:               transactionProto.Otr,
		AdminFee:          transactionProto.AdminFee,
		InstallmentAmount: transactionProto.InstallmentAmount,
		InterestAmount:    transactionProto.InterestAmount,
		AssetName:         transactionProto.AssetName,
		CreatedAt:         timestamppb.Now().AsTime(),
		UpdatedAt:         timestamppb.Now().AsTime(),
	}
	return r.db.WithContext(ctx).Table("transactions").Create(&transaction).Error
}

func (r *Repository) UpdateTransaction(ctx context.Context, transactionProto *transaction_proto.UpdateTransactionRequest) error {
	transactionDB, err := r.GetTransaction(ctx, transactionProto.ContractNumber)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(transactionDB.Id)
	if err != nil {
		return err
	}

	transaction := models.Transaction{
		Id:                uint64(id),
		CustomerId:        transactionProto.CustomerId,
		ContractNumber:    transactionProto.ContractNumber,
		Otr:               transactionProto.Otr,
		AdminFee:          transactionProto.AdminFee,
		InstallmentAmount: transactionProto.InstallmentAmount,
		InterestAmount:    transactionProto.InterestAmount,
		AssetName:         transactionProto.AssetName,
		UpdatedAt:         timestamppb.Now().AsTime(),
	}

	result := r.db.WithContext(ctx).Table("transactions").Save(&transaction)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *Repository) DeleteTransaction(ctx context.Context, contractNumber string) error {
	transactionDB, err := r.GetTransaction(ctx, contractNumber)
	if err != nil {
		return err
	}

	transactionId, err := strconv.Atoi(transactionDB.Id)
	if err != nil {
		return err
	}

	transaction := models.Transaction{
		Id:                uint64(transactionId),
		CustomerId:        transactionDB.CustomerId,
		ContractNumber:    transactionDB.ContractNumber,
		Otr:               transactionDB.Otr,
		AdminFee:          transactionDB.AdminFee,
		InstallmentAmount: transactionDB.InstallmentAmount,
		InterestAmount:    transactionDB.InterestAmount,
		AssetName:         transactionDB.AssetName,
	}

	result := r.db.WithContext(ctx).Table("transactions").Where("contract_number = ?", contractNumber).Delete(&transaction)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) Transaction(ctx context.Context, fn func(repo interfaces.TransactionRepository) error) error {
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
