
generate-proto:
	protoc \
		-I. \
		-I../sdk-helper/third_party \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=. --openapiv2_opt=logtostderr=true \
		cmd/proto/university.proto

start:
	docker-compose -f docker-compose.yaml up --build

run:
	docker-compose build && docker-compose up

stop:
	docker-compose -f docker-compose.yml down



.PHONY: up run down generate-proto