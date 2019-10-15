FROM golang:1.12.10 as build-env

LABEL maintainer="Damir Asylkhan<krosava_15d@mail.ru>"

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "test", "-v", "./..."]