BUILD_DIR := bin
TOOLS_DIR := tools

default: all

all: clean lint test build run

.PHONY: $(BUILD_DIR)/gateway
bin/gateway: cmd/gateway/*.go
	GOMEMLIMIT=100 GOGC=100 go build -ldflags="-s -w" -o ./bin/gateway -v cmd/gateway/main.go

.PHONY: $(BUILD_DIR)/user
bin/user: cmd/user/*.go
	GOMEMLIMIT=100 GOGC=100 go build -ldflags="-s -w" -o ./bin/user -v cmd/user/main.go

.PHONY: $(BUILD_DIR)/product
bin/product: cmd/product/*.go
	CGO_ENABLED=0 GOMEMLIMIT=100 GOGC=100 go build -ldflags="-s -w" -o ./bin/product -v cmd/product/main.go

.PHONY: build
build: bin/gateway bin/user bin/product

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod vendor
	@go mod tidy

.PHONY: run
run: build
	bin/gateway & bin/user & bin/product

.PHONY: hotreload
hotreload: $(TOOLS_DIR)/air/air
	./$(TOOLS_DIR)/air/air -c .air.gateway.toml & ./$(TOOLS_DIR)/air/air -c .air.user.toml & ./$(TOOLS_DIR)/air/air -c .air.product.toml

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

tools/cosmtrek/air:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh| sh -s -- -b tools/air latest

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run -v

.PHONY: test
test:
	go test -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

tools/protobuf/protoc-gen-go:
	mkdir -p tools/
	GOBIN=$(shell pwd)/tools go install github.com/golang/protobuf/protoc-gen-go

.PHONY: productProtoGateway
productProtoGateway:
	protoc -I="." --grpc-gateway_out ./ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true internal/adapter/api/v1/product/product.proto

.PHONY: productProto
productProto:
	protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. internal/adapter/api/v1/product/product.proto

.PHONY: userProtoGateway
userProtoGateway:
	protoc -I="." --grpc-gateway_out ./ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true internal/adapter/api/v1/user/user.proto

.PHONY: userProto
userProto:
	protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. internal/adapter/api/v1/user/user.proto

.PHONY: domainMessageProto
domainMessageProto:
	protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. internal/adapter/api/domain/proto/message.proto

.PHONY: domainErrorProto
domainErrorProto:
	protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. internal/adapter/api/domain/proto/error.proto

.PHONY: proto
proto: productProtoGateway productProto userProtoGateway userProto domainMessageProto domainErrorProto

.PHONY: openapi
openapi:
	chmod +x ./api/generateOpenAPIYaml.sh
	./api/generateOpenAPIYaml.sh

.PHONY: kill
kill:
	lsof -i tcp:8000 | awk 'NR!=1 {print $2}' | xargs kill
	lsof -i tcp:8001 | awk 'NR!=1 {print $2}' | xargs kill
	lsof -i tcp:8002 | awk 'NR!=1 {print $2}' | xargs kill
