FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o schedule ./cmd/server.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/schedule .
CMD ["./schedule"]
