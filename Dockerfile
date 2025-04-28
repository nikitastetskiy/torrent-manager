# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o torrent-manager ./cmd/api

# Final image
FROM gcr.io/distroless/static

COPY --from=builder /app/torrent-manager /torrent-manager

EXPOSE 8080

ENTRYPOINT ["/torrent-manager"]