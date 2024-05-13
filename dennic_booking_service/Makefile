-include .env
export

CURRENT_DIR=$(shell pwd)
APP=Booking_service
CMD_DIR=./cmd

.DEFAULT_GOAL = build
#POSTGRES_USER = postgres
#POSTGRES_PASSWORD = 20030505
#POSTGRES_HOST = localhost
#POSTGRES_PORT = 5432
#POSTGRES_DATABASE = dennic_booking_service

# run service
.PHONY: run
run:
	go run ${CMD_DIR}/app/main.go

.PHONY: create-migration
create-migration:
	migrate create -dir migrations -ext sql -seq $(name)_table

# proto-gen
.PHONY: proto-gen
proto-gen:
	./scripts/genproto.sh

DB_URL := "postgres://postgres:20030505@localhost:5432/dennic_booking_service?sslmode=disable"

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

.PHONY: migrate-force
migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1