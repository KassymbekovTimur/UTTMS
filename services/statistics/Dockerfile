FROM golang:1.24-alpine AS builder
# adding dependencies with C-modules jic
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN apk add --no-cache protoc protobuf protoc-gen-go protoc-gen-go-grpc && \
    protoc --proto_path=./proto \
      --go_out=paths=source_relative:./proto \
      --go-grpc_out=paths=source_relative:./proto \
      proto/stats.proto

RUN go build -o stats-service ./cmd/server

FROM alpine:3.18
#certificates for HTTPS/gRPC jic
RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/stats-service .

ENTRYPOINT ["./stats-service"]