package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/delivery"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/repository"
	"github.com/depri11/technical_kreditplus/customer_service/internal/app/usecase"
	"github.com/depri11/technical_kreditplus/customer_service/internal/pkg/db"
	"github.com/depri11/technical_kreditplus/customer_service/internal/pkg/opentelemetry"
	customer_proto "github.com/depri11/technical_kreditplus/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()

	db, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	tp, err := opentelemetry.NewTracerProvider(ctx, cfg.APP.ServiceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	port := fmt.Sprintf(":%d", cfg.APP.Port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	repo := repository.NewRepository(db, cfg)
	usecase := usecase.NewUseCase(repo, cfg)
	delivery := delivery.NewDelivery(usecase)

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Timeout:           15 * time.Second,
		MaxConnectionAge:  5 * time.Minute,
		Time:              5 * time.Minute,
	}))

	customer_proto.RegisterCustomerServiceServer(server, delivery)

	go func() {
		err := server.Serve(l)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Printf("GRPC Server Article is listening on port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server GRPC Article Exited Properly")
}
