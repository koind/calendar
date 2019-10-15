FROM golang:1.12.10 as build-env

LABEL maintainer="Damir Asylkhan<krosava_15d@mail.ru>"

WORKDIR /scheduler

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /opt/scheduler

# Release
FROM alpine:latest

WORKDIR /root/

COPY --from=build-env /opt/scheduler .
COPY --from=build-env /scheduler/config ./config

CMD ["./scheduler", "--config=./config/development/config.toml"]