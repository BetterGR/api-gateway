# Docker settings
DOCKER_IMAGE_NAME ?= $(notdir $(CURDIR))
DOCKER_TAG ?= latest
DOCKERFILE ?= Dockerfile
DOCKER_REGISTRY ?= ghcr.io/bettergr

# Default target
all: gomod fmt vet lint

# Ensure tools are installed
ensure-gofumpt:
ifeq ($(OS),Windows_NT)
	@where gofumpt > temp.txt 2>&1 || ( \
		echo [INSTALL] gofumpt not found. Installing... & \
		go install mvdan.cc/gofumpt@latest \
	)
	@ del temp.txt
else
	@command -v gofumpt > /dev/null 2>&1 || { \
		echo "[INSTALL] gofumpt not found. Installing..."; \
		go install mvdan.cc/gofumpt@latest; \
	}
endif

ensure-gci:
ifeq ($(OS),Windows_NT)
	@where gci > temp.txt 2>&1 || ( \
		echo [INSTALL] gci not found. Installing... & \
		go install github.com/daixiang0/gci@latest \
	)
	@ del temp.txt
else
	@command -v gci > /dev/null 2>&1 || { \
		echo "[INSTALL] gci not found. Installing..."; \
		go install github.com/daixiang0/gci@latest; \
	}
endif

ensure-golangci-lint:
ifeq ($(OS),Windows_NT)
	@where golangci-lint > temp.txt 2>&1 || ( \
		echo [INSTALL] golangci-lint not found. Installing... & \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
	)
	@ del temp.txt
else
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "[INSTALL] golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	}
endif

# Manage Go modules
gomod:
	@echo [GO-MOD] Verifying modules...
	@go mod tidy
	@go mod verify
	@echo [GO-MOD] Modules verified.

# Ensure gqlgen is installed
ensure-gqlgen:
ifeq ($(OS),Windows_NT)
	@where gqlgen > temp.txt 2>&1 || ( \
		echo [INSTALL] gqlgen not found. Installing... & \
		go install github.com/99designs/gqlgen@latest \
	)
	@ del temp.txt
else
	@command -v gqlgen > /dev/null 2>&1 || { \
		echo "[INSTALL] gqlgen not found. Installing..."; \
		go install github.com/99designs/gqlgen@latest; \
	}
endif

# Generate GraphQL code from schema
gqlgen: ensure-gqlgen
	@echo [GQLGEN] Generating GraphQL code from schema...
	@gqlgen generate
	@echo [GQLGEN] GraphQL code generated successfully.

# Format Go code
fmt: ensure-gofumpt ensure-gci
	@echo [FMT] Formatting Go code...
	@go fmt ./...
	@gofumpt -w .
	@gci write --skip-generated .
	@echo [FMT] Go code formatted.

# Vet Go code
vet:
	@echo [VET] Running vet checks on Go code...
	@go vet ./...
	@echo [VET] Vet checks completed.

# Lint Go code
lint: ensure-golangci-lint fmt
	@echo [LINT] Running linter on Go code...
	@golangci-lint run
	@echo [LINT] Lint checks completed.

# Build server
build: fmt vet lint
	@echo [BUILD] Building server binary...
	@go build -o server ./server.go
	@echo [BUILD] Server binary built successfully.

# Run the server
run: fmt vet
	@echo [RUN] Starting server...
	@go run ./server.go $(ARGS)

test: gomod fmt vet lint
	@echo [TEST] Running tests...
	@go test -v ./server/ | grep -v '=== RUN' | sed 's/--- PASS:/ [PASS]/' | sed 's/--- FAIL:/ [FAIL]/'
	@echo [TEST] Tests completed.

# Build Docker image
docker-build: gomod fmt vet lint build
	@echo [DOCKER] Building Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)... 
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) -f $(DOCKERFILE) .
	@echo [DOCKER] Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) built successfully.

# Push Docker image to registry
docker-push: docker-build
ifeq ($(DOCKER_REGISTRY),docker.io)
	@echo [DOCKER] Docker registry is set to docker.io.
else
	@echo [DOCKER] Docker registry set to $(DOCKER_REGISTRY).
endif
	@echo [DOCKER] Pushing Docker image $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) to $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)...
	@docker tag $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
	@docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
	@echo [DOCKER] Docker image pushed successfully.

# Clean up generated files
clean:
	@echo [CLEAN] Removing generated files...
ifeq ($(OS),Windows_NT)
	@del /Q server\server
	@del /Q protos\*.pb.go
else
	@rm -rf server/server
	@rm -rf protos/*.pb.go
endif
	@echo [CLEAN] Clean up complete.

help:
	@echo Available targets:
	@echo   all               Build and check everything (proto, gomod, fmt, vet, lint)
	@echo   proto             Generate Go code from proto file
	@echo   gomod             Manage Go modules (tidy and verify)
	@echo   gqlgen            Generate GraphQL code from schema
	@echo   fmt               Format Go code
	@echo   vet               Run vet checks on Go code
	@echo   lint              Run linter on Go code
	@echo   build             Build the server binary
	@echo   run               Run the server
	@echo   docker-build      Build Docker image
	@echo   docker-push       Push Docker image to registry
	@echo   clean             Clean up generated files

.PHONY: all proto fmt run vet lint build docker-build docker-push gomod clean ensure-gofumpt ensure-gci ensure-golangci-lint ensure-gqlgen gqlgen help test