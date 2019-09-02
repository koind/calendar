FROM golang:1.12

LABEL maintainer="Asylkhan Damir <assylkhan.d@mail.ru>"

ENV GO111MODULE=on

WORKDIR /calendar

COPY . .

RUN go mod download

RUN CALENDAR_HOST="localhost" CALENDAR_PORT=7777 go run /calendar/cmd/server/http_server.go