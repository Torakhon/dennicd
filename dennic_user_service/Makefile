-include .env
export

CURRENT_DIR=$(shell pwd)
APP=dennic_user_service
CMD_DIR=./cmd

.DEFAULT_GOAL = build
#POSTGRES_USER = postgres
#POSTGRES_PASSWORD = 20030505
#POSTGRES_HOST = localhost
#POSTGRES_PORT = 5432
#POSTGRES_DATABASE = dennic

# build for current os
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# run service
.PHONY: run
run:
	go run cmd/app/main.go

<<<<<<< HEAD
## migrate
#.PHONY: migrate
#migrate:
#	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up
#
#migrate-file:
#	migrate create -ext sql -dir migrations/ -seq create_table_users
#
## go generate
=======
# migrate
.PHONY: migrate-up
migrate-up:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

# migrate
.PHONY: migrate-down
migrate-down:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_table_users

# proto
.PHONY: proto-gen
proto-gen:
	./scripts/gen-proto.sh

# git submodule init 	
.PHONY: pull-proto
pull-proto:
	git submodule update --init --recursive

# go generate	
>>>>>>> 2e1c1519a2a2a6f1158dd26816fa8f74e469f23b
.PHONY: go-gen
#go-gen:
#	go generate ./...
#
## run test
#test:
#	go test -v -cover -race ./internal/...

# -------------- for deploy --------------
#build-image:
#	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
#	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}
#
#push-image:
#	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
#	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

.PHONY: pull-proto-module
pull-proto-module:
	git submodule update --init --recursive

.PHONY: update-proto-module
update-proto-module:
	git submodule update --remote --merge


.PHONY: proto-gen
proto-gen:
	./scripts/genproto.sh

DB_URL := "postgres://postgres:20030505@localhost:5432/dennic_user_service?sslmode=disable"

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
