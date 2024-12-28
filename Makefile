proto:
	@echo Generating article microservice proto
	protoc --go_out=protos --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=protos protos/*.proto

api:
	go run api_gateway/cmd/main.go

customer:
	go run customer_service/cmd/main.go

transaction:
	go run transaction_service/cmd/main.go

test:
	go test -v ./customer_service/testing