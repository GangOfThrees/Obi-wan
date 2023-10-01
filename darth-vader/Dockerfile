# syntax=docker/dockerfile:1

FROM golang:1.20 AS build

WORKDIR /app
COPY darth-vader/go.mod darth-vader/go.sum ./darth-vader/
COPY light-sabers/ ./light-sabers/

RUN ls -la /app/darth-vader
RUN ls -la /app

WORKDIR /app/darth-vader
RUN go mod download

COPY darth-vader/ ./

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o darth-vader ./cmd/server.go

FROM alpine:3.9.6

WORKDIR /app
COPY --from=build /app/darth-vader /app/darth-vader/.env /app/

RUN apk -U upgrade
RUN apk add --no-cache dumb-init ca-certificates
RUN chmod +x /app/darth-vader

EXPOSE 8001

ENTRYPOINT ["/app/darth-vader"]
