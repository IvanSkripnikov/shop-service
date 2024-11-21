FROM golang:1.20-alpine

WORKDIR /go/src/app

RUN apk update && \
    apk add git && \
    apk add --no-cache alpine-conf

ADD ./app /go/src/app
ADD ./app/log /go/log

RUN cd /go/src/app && \
    go install github.com/IvanSkripnikov/loyalty_system

WORKDIR "/go"

ENTRYPOINT loyalty_system