FROM golang:1.26.4 AS builder

WORKDIR /app

COPY warehouse/go.mod warehouse/go.sum ./
COPY gen/go /gen/go

RUN  go mod download


COPY warehouse/. ./ 

RUN CGO_ENABLED=0 GOOS=linux go build -o warehouse-service ./cmd/server

# run again
FROM alpine:latest

WORKDIR /app

COPY --from=builder app/warehouse-service .

CMD ["./warehouse-service","--warehouse-config=config/config.yaml"]