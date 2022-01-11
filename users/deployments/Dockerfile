FROM golang:1.17

LABEL maintaner="alexdenkk@yandex.ru"

RUN go mod download

RUN go build -o app cmd/main.go