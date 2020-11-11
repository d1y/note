# Author: d1y<chenhonzhou@gmail.com>

FROM golang:1.15.4-alpine3.12 as builder

WORKDIR /go/src/github.com/d1y/note

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct \
    PORT=2333

ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache gcc

# https://stackoverflow.com/q/36575471
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache sqlite
RUN apk add --no-cache build-base

RUN go get && make build_prod

FROM alpine:latest as prod

WORKDIR /go/note

COPY --from=0 /go/src/github.com/d1y/note/build /go/note

ENV APP_ENV=release

CMD [ "./note.app" ]

EXPOSE $PORT