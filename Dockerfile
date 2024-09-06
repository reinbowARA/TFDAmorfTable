FROM golang:1.22.4 AS base

FROM base AS dev

WORKDIR /code

COPY . .
RUN go mod tidy
RUN go run main.go
