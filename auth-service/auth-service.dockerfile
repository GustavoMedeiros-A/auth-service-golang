# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first and download modules for caching
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
# Install CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon


# CompileDaemon will run the application and watch for file changes
ENTRYPOINT CompileDaemon -log-prefix=false --build="go build -o authApp ./" --command=./authApp