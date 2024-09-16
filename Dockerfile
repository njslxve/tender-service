FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/tender-service ./cmd/tender-service/main.go

FROM alpine:latest

COPY --from=builder /app/bin/tender-service /tender-service

CMD ["/tender-service"]