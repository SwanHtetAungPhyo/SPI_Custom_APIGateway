# Build Stage
FROM golang:1.23.2-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the entire project to the working directory
COPY . .

# Install necessary dependencies (if needed)
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Run Stage (using a smaller base image)
FROM debian:bullseye-slim

# Install necessary packages (e.g., for running Go binary)
RUN apt-get update && apt-get install -y ca-certificates

# Set the current working directory inside the container
WORKDIR /root/

# Copy the Go binary from the build stage
COPY --from=build /app/main .

# Expose the application port
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
