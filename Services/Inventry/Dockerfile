FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o inventory-service ./cmd/server

CMD ["./inventory-service","--inventory-config=config/config.yaml"]