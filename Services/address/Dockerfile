FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o address-service ./cmd/server

CMD ["./address-service","--address-config=config/config.yaml"]