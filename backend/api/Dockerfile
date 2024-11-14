# Use the official Golang image as the build stage
FROM golang:1.19.2-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies to cache them in a separate layer
RUN go mod vendor

# Copy the source code into the container
COPY . .

# Build the Go application with optimizations
RUN go build -o gate main.go

# Use a minimal base image for the final stage
FROM debian:bullseye-slim

# Copy the built application from the builder stage
COPY --from=builder /app/gate /gate

# Set up the default port
EXPOSE 8082

# Set the entrypoint to the built binary
CMD ["/gate"]
