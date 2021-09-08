FROM golang:1.16.4-alpine

ENV TERM=xterm
ENV TZ=Asia/Bangkok

RUN mkdir -p /usr/app/src
WORKDIR /usr/app/src

COPY ./src/.air.toml .
COPY ./src .

RUN go get -u github.com/cosmtrek/air
RUN go mod tidy
RUN go mod download