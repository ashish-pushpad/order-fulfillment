FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o checkout-service ./cmd/server

CMD ["./checkout-service","--checkout-config=config/config.yaml"]