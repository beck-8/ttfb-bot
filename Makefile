.PHONY: all build gen clean run test

APP_NAME := ttfb
ABI_DIR := contracts
GEN_DIR := pkg/contracts

all: build

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) cmd/ttfb/main.go

gen:
	@echo "Generating contract bindings..."
	@mkdir -p $(GEN_DIR)/registry $(GEN_DIR)/warmstorage $(GEN_DIR)/view $(GEN_DIR)/verifier
	@abigen --abi $(ABI_DIR)/ServiceProviderRegistry.abi --pkg registry --type ServiceProviderRegistry --out $(GEN_DIR)/registry/registry.go
	@abigen --abi $(ABI_DIR)/WarmStorageService.abi --pkg warmstorage --type WarmStorageService --out $(GEN_DIR)/warmstorage/warmstorage.go
	@abigen --abi $(ABI_DIR)/WarmStorageView.abi --pkg view --type WarmStorageView --out $(GEN_DIR)/view/view.go
	@abigen --abi $(ABI_DIR)/PDPVerifier.abi --pkg verifier --type PDPVerifier --out $(GEN_DIR)/verifier/verifier.go
	@echo "Done."

clean:
	@echo "Cleaning..."
	@rm -f $(APP_NAME)
	@rm -rf bin/
	@rm -rf $(GEN_DIR)/*/*.go

run: build
	@./$(APP_NAME)

test:
	@go test ./...
