FROM golang:1.20.10-alpine3.17 as Builder

RUN apk add --update  && \
    apk add --no-cache alpine-conf tzdata git

RUN git clone https://github.com/IvanSkripnikov/loyalty_system.git && \
    cd loyalty_system && \
    git checkout main

ADD ./loyalty_system/src /go/src/loyalty_system
ADD ./loyalty_system/src/config /go/config

RUN cd /go/src/loyalty_system && \
    go install loyalty_system

FROM alpine:3.18.4 as App

COPY --from=Builder /go/bin/* /go/bin/loyalty_system
COPY --from=Builder /go/config /go/config

EXPOSE 8080

WORKDIR "/go"
ENTRYPOINT ["/go/bin/loyalty_system"]