package testing_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/depri11/technical_kreditplus/transaction_service/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/testcontainers/testcontainers-go"
	psql "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var r *mux.Router
var db *gorm.DB

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cfg := &config.Config{
		APP: config.App{
			ServiceName: "transaction_service",
			Port:        8080,
		},
		DB: config.DB{
			User:     "username",
			Password: "password",
			Host:     "localhost",
			Port:     5432,
			Name:     "transaction_service",
		},
	}

	postgresContainer, err := psql.Run(ctx,
		"postgres:latest",
		psql.WithDatabase(cfg.DB.Name),
		psql.WithUsername(cfg.DB.User),
		psql.WithPassword(cfg.DB.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		panic(err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, host, port.Port(), cfg.DB.Name)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db = gormDB

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	err = MigrationDB(ctx, cfg, dsn, true)
	if err != nil {
		panic(err)
	}
	defer MigrationDB(ctx, cfg, dsn, false)
	m.Run()
}

func MigrationDB(ctx context.Context, cfg *config.Config, dsn string, isUp bool) error {
	// cwd, _ := os.Getwd()
	m, err := migrate.New("file://../migrations", dsn)
	if err != nil {
		return err
	}

	if isUp {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return err
		}
	} else {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}
