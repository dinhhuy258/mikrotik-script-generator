FROM golang:1.23.2-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ../go.mod ../go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code to the working directory
COPY ../ ./

# Build the Go app
RUN go build -o /app/mikrotik-script-generator main.go

ENTRYPOINT ["/app/mikrotik-script-generator", "service"]
