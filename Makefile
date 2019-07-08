NAME := kafka-prompt

.DEFAULT_GOAL := all

.PHONY: gen-setting
gen-setting: ## Generate setting source code from json config
	mkdir -p ./bin
	go build -o ./bin/generator ./settings/generator
	./bin/generator settings/kafka-command.json kaprompt/suggests.go

.PHONY: build
build: main.go
	GO111MODULE=on go build -o bin/kaprompt

.PHONY: all
all: gen-setting build

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'

