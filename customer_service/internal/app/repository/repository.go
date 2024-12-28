package repository

import (
	"context"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
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

func (repo *repository) GetCustomer(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error) {
	return nil, nil
}
