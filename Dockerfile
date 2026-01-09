# Use Go base image for building
FROM golang:latest AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory for api-gateway
WORKDIR /app

# Copy go.mod and go.sum files
COPY ./go.mod ./go.sum ./

COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o main ./cmd/main.go

# Final stage
FROM alpine:3.20

# Set working directory
WORKDIR /root/

# Install certificates
RUN apk --no-cache add ca-certificates

# Create directory for .env file
RUN mkdir -p /root/internal

# Copy built binary from builder to a specific file path
COPY --from=builder /app /root/

# # Copy .env file
# COPY ./.env /root/.env

# Expose port defined by REST_PORT environment variable
EXPOSE $REST_PORT

# Run the app
CMD ["./main"]
