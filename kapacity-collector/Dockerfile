FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download
RUN go mod vendor 
RUN go mod tidy

# Build the application and set executable permissions
RUN go build -o /kapacity-collector

RUN chmod +x /kapacity-collector

EXPOSE 8080

CMD ["/kapacity-collector"]

