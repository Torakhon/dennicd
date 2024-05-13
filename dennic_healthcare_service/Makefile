CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: proto-gen
proto-gen:
	./scripts/genproto.sh

DB_URL := "postgres://postgres:20030505@localhost:5432/dennic_healthcare_service?sslmode=disable"

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

.PHONY: migrate-force
migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1

.PHONY: migrate-file
migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_doctorsService_table


.PHONY: pull-proto-module
pull-proto-module:
	git submodule update --init --recursive

.PHONY: update-proto-module
update-proto-module:
	git submodule update --remote --merge
