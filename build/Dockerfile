FROM golang:1.23.2-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY mikrotik-script-generator/go.mod mikrotik-script-generator/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY mikrotik-script-generator/. .

# Build the Go app
RUN go build -o main main.go

# Command to run the executable
ENTRYPOINT ["./main", "service"]
