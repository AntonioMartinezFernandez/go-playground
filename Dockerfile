FROM golang:1.22.6-alpine3.20 as BUILDER
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o cmd/main.go

FROM scratch
COPY --from=BUILDER /app/app /app
ENTRYPOINT ["/app"]