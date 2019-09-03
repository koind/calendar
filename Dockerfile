FROM golang:1.12

LABEL maintainer="Asylkhan Damir <assylkhan.d@mail.ru>"

#ENV GO111MODULE=on

WORKDIR /calendar

COPY . .

EXPOSE 7777

RUN go mod download