# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o main ./app/main.go

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /database/database.sqlite ./database.sqlite

# Expose the port Gin will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
