package usecase

import (
	"context"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/interfaces"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type usecase struct {
	repoArticle interfaces.CustomerRepository
	tracer      trace.Tracer
}

func NewUseCase(repoArticle interfaces.CustomerRepository, cfg *config.Config) *usecase {
	tracer := otel.Tracer(cfg.APP.ServiceName)
	return &usecase{repoArticle, tracer}
}

func (u *usecase) GetCustomer(ctx context.Context, id string) (*customer_proto.GetCustomerResponse, error) {
	return nil, nil
}
