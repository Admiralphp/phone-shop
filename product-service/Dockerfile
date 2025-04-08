# Use the official Golang image to build the app
FROM golang:1.23-alpine as builder

# Set the environment variable for Go modules and GOPATH
ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Set the working directory inside the container
WORKDIR /go/src/product-service

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Install swag for generating swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate the swagger docs
RUN swag init

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o product-service ./main.go

# Start a new stage for the final image
FROM alpine:3.16

# Install the required dependencies in the final image
RUN apk --no-cache add ca-certificates tzdata

# Set the working directory for the app
WORKDIR /app

# Copy the binary and swagger docs from the builder stage
COPY --from=builder /go/src/product-service/product-service /app/
COPY --from=builder /go/src/product-service/docs /app/docs/

# Expose the required port (if needed)
EXPOSE 8080

# Command to run the application
CMD ["./product-service"]
