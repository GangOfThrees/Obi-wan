# syntax=docker/dockerfile:1

FROM golang:1.20 AS build

WORKDIR /app
COPY kenobi/go.mod kenobi/go.sum ./kenobi/
COPY light-sabers/ ./light-sabers/

WORKDIR /app/kenobi
RUN go mod download

COPY kenobi/ ./

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kenobi ./cmd/server.go

FROM alpine:3.9.6

WORKDIR /app
COPY --from=build /app/kenobi /app/kenobi/.env /app/

RUN apk -U upgrade
RUN apk add --no-cache dumb-init ca-certificates
RUN chmod +x /app/kenobi

EXPOSE 8001

ENTRYPOINT ["/app/kenobi"]
