# syntax=docker/dockerfile:1

FROM golang:1.20 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kenobi ./cmd/server.go

FROM alpine:3.9.6

WORKDIR /app
COPY --from=build /app/kenobi /app/.env /app/

RUN apk -U upgrade
RUN apk add --no-cache dumb-init ca-certificates
RUN chmod +x /app/kenobi

EXPOSE 8000

ENTRYPOINT ["/app/kenobi"]
