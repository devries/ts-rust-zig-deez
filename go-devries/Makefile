##
# Static part, don't change these:
##

.PHONY: help ready docker-build docker-ready fmt lint test

help:
	@cat Makefile | grep -E "^\w+$:"

ready: fmt lint test

docker-build:
	docker build --no-cache . -t deez_$(notdir $(shell pwd))

docker-ready: docker-build
	docker run -v $(shell pwd):/deez -t deez_$(notdir $(shell pwd))

##
# Update those:
##

clean:
	@echo "===> Cleaning"
	rm lexer/tokentype_string.go || true

fmt:
	@echo "===> Formatting"
	go fmt ./...

lint: lexer/tokentype_string.go
	@echo "===> Linting"
	go vet ./...

test: test-lexer test-ast test-parser
	@echo "===> Testing EVERYTHING"

test-lexer: lexer/tokentype_string.go
	@echo "===> Testing lexer"
	go test ./lexer
	
test-ast: lexer/tokentype_string.go
	@echo "===> Testing AST"
	go test ./ast

test-parser: lexer/tokentype_string.go
	@echo "===> Testing parser"
	go test ./parser
	
lexer/tokentype_string.go: lexer/lexer.go
	go generate monkey/lexer
