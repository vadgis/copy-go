FROM alpine:latest as builder

RUN apk update && \
 apk upgrade && \
 apk add go git make 

COPY / /tmp/builder

WORKDIR /tmp/builder

ENTRYPOINT [ "/bin/ash" ]