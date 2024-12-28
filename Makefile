init: proto
	go mod tidy

proto:
	@echo Generating article microservice proto
	protoc --go_out=protos --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=protos protos/*.proto

customer_service:
	go run cmd/customer-service/main.go

transaction_service:
	go run cmd/transaction-service/main.go