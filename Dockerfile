# commands:
# docker build -t meditrack-api .

# Start from this golang base image
FROM golang:1.23.4-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /

# Copy go mod and sum files and download dependencies before copying the rest of the files. This is
# done to take advantage of caching and avoid downloading dependencies every time the code changes
COPY go.mod go.sum ./
RUN go mod download

# Copy go mod and sum files
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Start from scratch image
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /main /

# ENV DATABASE_URL=$DATABASE_URL \
#     OUTER_API_URL=$OUTER_API_URL \
#     OUTER_API_TOKEN=$OUTER_API_TOKEN \
#     OUTER_API_AUTH_HEADER=$OUTER_API_AUTH_HEADER

# Command to run the executable
ENTRYPOINT [ "/main" ]