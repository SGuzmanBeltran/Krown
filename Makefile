include .env
export

PROTO_NAME ?= user

run-user:
	@go run cmd/user/main.go

show-env:
	@echo $$DB_URL

db-status:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose status

db-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose up

db-reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=${GOOSE_DBSTRING} GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} goose reset

proto-gen:
	@protoc \
		--proto_path=protobuf "protobuf/$(PROTO_NAME).proto" \
		--go_out=services/genproto/$(PROTO_NAME) --go_opt=paths=source_relative \
		--go-grpc_out=services/genproto/$(PROTO_NAME) --go-grpc_opt=paths=source_relative