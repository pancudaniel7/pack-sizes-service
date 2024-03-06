# Start from the official Golang image, which includes the Go toolchain
FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY go.mod go.sum ./

RUN sed -i 'N s/go 1.21.5/go 1.23/' /app/go.mod
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/pack-sizes-service/main.go"]
