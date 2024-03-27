# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

ENV GO111MODULE=on

COPY . .
RUN go install

# Build
RUN go build -o /api

EXPOSE 8001

# Run
CMD ["/api"]
