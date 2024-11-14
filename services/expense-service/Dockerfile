# Use Golang image to build the application
FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o app .

# Use an Alpine image for the final stage (smaller image)
FROM alpine:latest

# Install certificates (needed for Go applications)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the app will run on
EXPOSE 4001 4002 4003

# Command to run the executable
CMD ["./app"]
