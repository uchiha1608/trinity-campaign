# Use an official Golang image as the base image
FROM golang:1.21-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files into the container (for dependency management)
COPY go.mod go.sum ./

# Download dependencies based on the go.mod and go.sum files
RUN go mod tidy

# Copy the rest of the application source code into the container
COPY . .

# Build the Go app
RUN go build -o trinity-campaign .

# Use a smaller image to run the application
FROM alpine:latest

# Install necessary packages (e.g., for MongoDB connectivity, CA certificates for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage into the final image
COPY --from=builder /app/trinity-campaign .

# Expose the port that your Go application will listen on
EXPOSE 8080

# Command to run your Go application
CMD ["./trinity-campaign"]
