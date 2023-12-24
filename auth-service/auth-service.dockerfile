# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy all project
COPY . . 

# Run go mod tidy to synchronize and verify dependencies
RUN go mod tidy

# Download necessary dependencies
RUN go mod download

# Build the Go application
RUN go build -o authApp ./cmd

# Expose port if your application runs on a specific port
# EXPOSE 80

# Set the entrypoint command to start the application
CMD ["./authApp"]