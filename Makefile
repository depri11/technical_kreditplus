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