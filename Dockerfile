FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY . .

RUN go mod download
RUN go mod vendor 
RUN go mod tidy

RUN go build -o /krv

EXPOSE 80

CMD [ "/krv" ]
