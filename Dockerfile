FROM ubuntu:bionic
FROM golang:1.22

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    gnupg

RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
    apt-get install -y migrate

CMD ["migrate -path /app/schema -database postgres://postgres:postgres@db:5432/postgres?sslmode=disable up"]

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN env GOOS=linux GOARCH=amd64 go build -o start core/main.go

EXPOSE 8000

CMD ["./start"]
