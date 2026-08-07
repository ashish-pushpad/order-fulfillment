FROM golang:1.26.4 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o product-service ./cmd/server

CMD ["./product-service", "--product-config=config/config.yaml"]