# syntax=docker/dockerfile:1

FROM golang:1.19
LABEL authors="mhuseinj"

WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./

RUN go build -o /bankTransfer
EXPOSE 8001

CMD ["/bankTransfer"]
