APP_NAME := govnoed-api
BIN_DIR := bin
CMD_PATH := ./cmd/api

.PHONY: build run check clean

build:
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

check:
	go vet ./...

deploy:
	go vet ./...
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)
	sudo systemctl restart govnoed-api