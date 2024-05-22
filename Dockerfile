FROM golang:1.22-alpine as build

ENV GOOS=linux
ENV GOARCH=amd64 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
RUN go mod vendor 
RUN go mod tidy

# Build the application and set executable permissions
RUN go build -o /kube-resource-viewer

EXPOSE 8080

CMD ["/kube-resource-viewer"]
