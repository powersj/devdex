.DEFAULT_TARGET: all

.PHONY: all
all: clean lint test build run

.PHONY: build
build:
	wails build

.PHONY: clean
clean:
	rm -f build/bin/devdex coverage.out
	rm -rf frontend/.svelte-kit frontend/build frontend/node_modules frontend/src/lib/wailsjs

.PHONY: deps
deps:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: dev
dev:
	wails dev

.PHONY: help
help:
	@echo 'Available Targets:'
	@echo '  build    build via wails'
	@echo '  clean    delete the build and test artifacts'
	@echo '  deps     install development dependencies'
	@echo '  dev      run in development mode'
	@echo '  help     print this output'
	@echo '  lint     run golangci-lint'
	@echo '  run      build and run binary'
	@echo '  test     run all unit tests'

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run: build
ifeq ($(shell uname), Linux)
	./build/bin/devdex
endif
ifeq ($(shell uname), Darwin)
	./build/bin/devdex.app/Contents/MacOS/devdex
endif


.PHONY: test
test:
	go test ./...
