# Use an official Go runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files do not change
RUN go mod vendor
RUN go mod tidy

# Build the Go app
RUN go build -o /kube-resource-viewer

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the executable
CMD ["/kube-resource-viewer"]
