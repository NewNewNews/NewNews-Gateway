FROM golang:1.21.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o gateway ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/gateway .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./gateway"]