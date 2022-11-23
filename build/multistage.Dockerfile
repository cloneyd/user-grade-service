# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal

WORKDIR /app/cmd/server

RUN go build -o /user-grade-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

ENV BASIC_USERNAME=cloneyd
ENV BASIC_PASSWORD=hello
ENV BACKUP_PATH=/backup/

COPY --from=build /user-grade-service /user-grade-service
COPY ./config /config

USER nonroot:nonroot

ENTRYPOINT ["/user-grade-service"]