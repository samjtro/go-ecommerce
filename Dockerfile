# syntax=docker/dockerfile:1

FROM golang:1.18-bullseye
WORKDIR /app

COPY go.* ./
COPY *.go ./

RUN go mod download
RUN go build -o /api

EXPOSE 8080

CMD [ "/api" ]