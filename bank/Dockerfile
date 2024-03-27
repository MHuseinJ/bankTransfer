FROM golang:1.19

# Set destination for COPY
WORKDIR /app

ENV GO111MODULE=on

COPY . .
RUN go install

# Build
RUN go build -o /bank

EXPOSE 8002

# Run
CMD ["/bank"]
