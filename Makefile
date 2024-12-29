proto:
	@echo Generating gRPC proto
	protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. protos/*.proto

api:
	go run api_gateway/cmd/main.go

customer:
	go run customer_service/cmd/main.go

transaction:
	go run transaction_service/cmd/main.go

test:
	go test -v ./customer_service/testing
	go test -v ./transaction_service/testing

DB_SOURCE=postgresql://username:password@localhost:5432/$(DB_NAME)?sslmode=disable

migrate-customer:
	@echo "Migrating Customer database..."
	@migrate -path customer_service/migrations -database "postgresql://username:password@localhost:5431/customer_service?sslmode=disable" --verbose up 1
	@echo "Database migrated."

migrate-transaction:
	@echo "Migrating Transaction database..."
	@migrate -path transaction_service/migrations -database "postgresql://username:password@localhost:5432/transaction_service?sslmode=disable" --verbose up 1
	@echo "Database migrated."