package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/depri11/technical_kreditplus/api_gateway/config"
	customerDelivery "github.com/depri11/technical_kreditplus/api_gateway/internal/customer/delivery"
	customerUsecase "github.com/depri11/technical_kreditplus/api_gateway/internal/customer/usecase"
	transactionDelivery "github.com/depri11/technical_kreditplus/api_gateway/internal/transaction/delivery"
	transactionUsecase "github.com/depri11/technical_kreditplus/api_gateway/internal/transaction/usecase"
	"github.com/depri11/technical_kreditplus/protos"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type server struct {
	mux *mux.Router
	v   validator.Validate
}

func NewServer(mux *mux.Router) *server {
	return &server{mux: mux, v: *validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()

	customerServicePort := fmt.Sprintf(":%d", cfg.CUSTOMER_SERVICE.Port)
	customerServiceconn, err := grpc.DialContext(ctx, customerServicePort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer customerServiceconn.Close()

	customerClient := protos.NewCustomerServiceClient(customerServiceconn)

	transactionServicePort := fmt.Sprintf(":%d", cfg.TRANSACTION_SERVICE.Port)
	transactionServiceconn, err := grpc.DialContext(ctx, transactionServicePort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer transactionServiceconn.Close()

	transactionClient := protos.NewTransactionServiceClient(transactionServiceconn)

	csUsecase := customerUsecase.NewCustomerUseCase(customerClient)
	csDelivery := customerDelivery.NewDelivery(csUsecase, s.mux)
	csDelivery.Routes()

	trxUsecase := transactionUsecase.NewTransactionUseCase(transactionClient)
	trxDelivery := transactionDelivery.NewDelivery(trxUsecase, s.mux)
	trxDelivery.Routes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	port := fmt.Sprintf(":%d", cfg.APP.Port)

	go func() {
		server := &http.Server{
			Addr:         port,
			Handler:      s.mux,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		}

		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Println("API Gateway listen on port", port)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server API Gateway Exited Properly")

	return nil
}
