# Go environment settings
GO := go
GO_FLAGS := -mod=vendor
BIN_NAME = gestimate

# Build the Go binary
.PHONY: build
build:
	@echo "Building the Go application..."
	$(GO) build $(GO_FLAGS) -o $(BIN_NAME)

# Run the Go application
.PHONY: run
run: build
	@echo "Running the Go application..."
	./$(BIN_NAME)

# Clean the Go binary
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm ./$(BIN_NAME)

include Makefile.docker.mk