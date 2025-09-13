FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app .

FROM alpine

WORKDIR /opt

COPY --from=builder /usr/local/bin/app .

EXPOSE 2815

CMD ["./app"]
