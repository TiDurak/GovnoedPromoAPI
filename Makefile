APP_NAME := govnoed-api
BIN_DIR := bin
CMD_PATH := ./cmd/api

.PHONY: build run clean

build:
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

clean:
	rm -f $(BIN_DIR)/$(APP_NAME)