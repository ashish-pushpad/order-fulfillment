FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o warehouse-service ./cmd/server

CMD ["./warehouse-service","--warehouse-config=config/config.yaml"]