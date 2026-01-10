# Variables
DOCKER_GO_IMAGE_NAME=coderbillzay/timesheet-api
DOCKERFILE_GO_PATH=./Dockerfile
DOCKER_CONTEXT=.
DOCKER_COMPOSE_FILE=./docker-compose.yml
DOCKER_COMPOSE_CMD=docker-compose -f $(DOCKER_COMPOSE_FILE) 

# Default target
production: swagger build run-docker-detached

development: swagger run

development-docker: swagger build run-docker

build-image: swagger build

# Format & Update swagger docs folder
swagger:
	@echo "Generate Swagger files ..."
	swag fmt --dir ./cmd,.
	swag init --generalInfo main.go --dir ./cmd,.

run:
	@echo "Running Go server ..."
	go run ./cmd/main.go

build:
	@echo "Building Docker image ..."
	docker build -t $(DOCKER_GO_IMAGE_NAME) .

run-docker:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE_CMD) up 

run-docker-detached:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE_CMD) up -d
