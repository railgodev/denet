FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o denet-users-service ./cmd/denet

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/denet-users-service /app/denet-users-service

COPY --from=builder /app/.env /app/.env


EXPOSE 8080

CMD ["/app/denet-users-service"]