FROM golang:1.18.1-alpine3.15 AS build

WORKDIR /usr/src/app

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v

FROM alpine:3.15.4

COPY --from=0 "/usr/src/app/gin-helloworld" gin-helloworld

ENTRYPOINT ./gin-helloworld
