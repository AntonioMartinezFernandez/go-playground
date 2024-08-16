# Builder image
FROM golang:1.22.6-alpine3.20 AS builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/EXAMPLE/main.go

# Distribution image
FROM scratch
COPY --from=builder /app/app /app

EXPOSE 8080/tcp
ENTRYPOINT ["/app"]