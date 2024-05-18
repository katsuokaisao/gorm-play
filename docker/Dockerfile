FROM golang:1.20.5-alpine3.18

RUN apk add --no-cache \
    gcc \
    libc-dev \
    tzdata \
    ca-certificates

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN go install github.com/cortesi/modd/cmd/modd@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .
