include .env
export

PROJECT_NAME=$(shell basename "$(PWD)")

## run-go: Run go app
run-go:
	go run main.go

## docker-up-app: Run docker app
docker-up-app:
	docker compose up -d app

## docker-down: Docker compose down
docker-down:
	docker compose down

## migrate-up: Migration up to the top point
migrate-up:
	@migrate -database ${CONNECTION_PATH} -path migrations up

## migrate-up-v: Migration up with version
migrate-up-v:
	@migrate -database ${CONNECTION_PATH} -path migrations up ${v}

## migrate-down: Migration down to start point
migrate-down:
	@migrate -database ${CONNECTION_PATH} -path migrations down

## migrate-down-v: Migration down with version
migrate-down-v:
	@migrate -database ${CONNECTION_PATH} -path migrations down ${v}

## help: help command
help: Makefile
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'