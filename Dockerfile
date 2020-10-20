# Author: d1y<chenhonzhou@gmail.com>

FROM daocloud.io/library/golang:1.7.6-alpine3.5

WORKDIR /go/src/github.com/d1y/note

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct \
    PORT=2333

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

RUN go build -o /go/bin/note .

CMD [ "note" ]

EXPOSE 2333