FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gateway .

FROM alpine
WORKDIR /app
COPY --from=builder /app/gateway .
CMD ["./gateway"]
