# syntax=docker/dockerfile:1

FROM golang:1.19
LABEL authors="mhuseinj"

WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./

RUN go build -o /bankApi
EXPOSE 8002

CMD ["/bankApi"]
