# commands:
# docker build -t meditrack-api .

# Start from this golang base image
FROM golang:1.23.4-alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy go mod and sum files
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Build the Go app
RUN go build -o main cmd/main.go

# Command to run the executable
CMD ["./main"]