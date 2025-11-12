PHONY: default run

APP_NAME=WebNewsAPI

default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go
