# Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /rating-app-api

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Set the working directory
WORKDIR /rating-app-api/cmd

# Build the Go app
RUN go build -o /rating-app-api/cmd/main

# Use a minimal image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /rating-app-api

# Copy the binary from the builder image
COPY --from=builder /rating-app-api/cmd/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
