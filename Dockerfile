FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
RUN go mod vendor 
RUN go mod tidy

# Build the application and set executable permissions
RUN go build -o /krv && chmod +x /krv

EXPOSE 80

CMD ["/krv"]
