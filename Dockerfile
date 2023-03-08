FROM golang:1.18.1-buster

MAINTAINER Agu Norbert <agunorbert@gmail.com>

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app

EXPOSE 8080

CMD ["./app"]