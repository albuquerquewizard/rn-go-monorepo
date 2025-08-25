FROM golang:1.25-alpine

# Install git for go mod download
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the app
RUN go build -o main ./cmd/go-boilerplate/main.go

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]