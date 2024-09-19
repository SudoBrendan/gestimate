# Docker settings
DOCKER_IMAGE := $(BIN_NAME)
DOCKER_FILE := Dockerfile
DOCKER_TAG := latest

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) -f $(DOCKER_FILE) .

# Run the Docker container
.PHONY: docker-run
docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm $(DOCKER_IMAGE):$(DOCKER_TAG)