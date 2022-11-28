.PHONY: build clean subjects variables server

subjects:
	go run cmd/crawler/main.go --resource subjects
variables:
	go run cmd/crawler/main.go --resource variables
server:
	go run cmd/api/main.go
test:
	go test ./internal/... ./pkg/...
OPENAPI_GENERATOR_VERSION=v6.2.1
OPENAPI_GENERATOR=go
clean:
	rm -rf pkg/client_*
build: pkg/client_gus
	goimports -w pkg/*
pkg/client_gus:
	mkdir pkg/client_gus
	cp api/.openapi-generator-ignore pkg/client_gus/.openapi-generator-ignore
	docker run --rm -v "$(PWD):/local" openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
		-g $(OPENAPI_GENERATOR) \
		-c /local/api/config-gus.json \
		-i /local/api/openapi-gus.json \
		-o /local/pkg/client_gus \
		-t /local/api/tpl


ifneq (,$(wildcard ./.env))
    include .env
    export
endif
