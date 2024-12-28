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

	"github.com/depri11/technical_kreditplus/customer_service/config"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
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
