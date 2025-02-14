# commands:
# docker build -t meditrack-api .

# Start from this golang base image
FROM golang:1.23.4-alpine AS stage1

# Set the Current Working Directory inside the container
WORKDIR /

# Copy go mod and sum files
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Start from scratch image
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=stage1 /main /

# Command to run the executable
ENTRYPOINT [ "/main" ]