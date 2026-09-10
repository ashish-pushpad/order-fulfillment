FROM golang:1.26.4 AS builder

WORKDIR /app

COPY user/go.mod user/go.sum ./
# COPY go.mod go.sum ./
COPY gen/go /gen/go
RUN go mod download

COPY user/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

CMD ["./user-service","--user-config=config/config.yaml"]