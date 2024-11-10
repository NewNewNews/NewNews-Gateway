# Builder stage
FROM golang:1.21.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV DATABASE_URL=postgresql://user:password@db:5432/gateway?schema=public

COPY prisma ./prisma
RUN go run github.com/steebchen/prisma-client-go generate

COPY . .
RUN go build -o gateway ./cmd/server

# Runtime stage
FROM golang:1.21.6-alpine

WORKDIR /root/

# Install necessary tools
RUN apk add --no-cache wget postgresql-client

# Copy binary and necessary files from builder
COPY --from=builder /app/gateway .
COPY --from=builder /app/.env .
COPY --from=builder /app/prisma ./prisma

# Copy wait-for script and startup script
COPY wait-for.sh /usr/local/bin/wait-for
COPY startup.sh .

# Make scripts executable
RUN chmod +x /usr/local/bin/wait-for startup.sh

EXPOSE 8080

CMD ["./startup.sh"]