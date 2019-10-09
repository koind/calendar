FROM golang:1.12.10 as build-env

LABEL maintainer="Damir Asylkhan<krosava_15d@mail.ru>"

WORKDIR /notifier

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /opt/notifier

# Release
FROM alpine:latest

COPY --from=build-env /opt/notifier /bin/notifier

CMD ["/bin/notifier"]