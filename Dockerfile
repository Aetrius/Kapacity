FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN go mod download
RUN go mod vendor
RUN go mod tidy

RUN go build -o /kube-resource-viewer

EXPOSE 8080

CMD [ "/kube-resource-viewer" ]