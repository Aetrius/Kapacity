FROM golang:1.22-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go mod vendor
RUN go mod tidy

# Copy the source code into the container
COPY *.go ./

# Build the Go app
RUN go build -o /kube-resource-viewer

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/kube-resource-viewer"]
