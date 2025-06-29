# Dockerfile
FROM golang:1.24.3-alpine

# Set environment variables
ENV GO111MODULE=on

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN go build -o app .

EXPOSE 8080

# Set the default command to run the compiled binary
CMD ["./app"]