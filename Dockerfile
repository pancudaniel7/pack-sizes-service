# Start from the official Golang image, which includes the Go toolchain
FROM golang:1.22.1-alpine3.19

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/pack-sizes-service/main.go"]
