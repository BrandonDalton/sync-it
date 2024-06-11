FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENV GCP_BUCKET=""
ENV DIR=""

RUN go build -o sync-it

CMD ["./sync-it"]