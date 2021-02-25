FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /opt/golang-todo
COPY . .

RUN go mod download

RUN go get -u github.com/cosmtrek/air

ENTRYPOINT air