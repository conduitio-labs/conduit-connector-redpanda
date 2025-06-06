VERSION=$(shell git describe --tags --dirty --always)

.PHONY: build
build:
	go build -ldflags "-X 'github.com/conduitio-labs/conduit-connector-redpanda.version=${VERSION}'" -o conduit-connector-redpanda cmd/connector/main.go

.PHONY: test-kafka
test-kafka:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose-kafka.yml up --quiet-pull -d --wait
	go test $(GOTEST_FLAGS) -race ./...; ret=$$?; \
		docker compose -f test/docker-compose-kafka.yml down; \
		exit $$ret

.PHONY: test-redpanda
test-redpanda:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose-redpanda.yml up --quiet-pull -d --wait
	go test $(GOTEST_FLAGS) -race ./...; ret=$$?; \
		docker compose -f test/docker-compose-redpanda.yml down; \
		exit $$ret

.PHONY: test
test: test-kafka test-redpanda

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt
fmt: ## Format Go files using gofumpt and gci.
	gofumpt -l -w .
	gci write --skip-generated  .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/go.mod
	@go list -modfile=tools/go.mod tool | xargs -I % go list -modfile=tools/go.mod -f "%@{{.Module.Version}}" % | xargs -tI % go install %
	@go mod tidy
