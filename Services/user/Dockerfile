FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o user-service ./cmd/server

CMD ["./user-service","--user-config=config/config.yaml"]