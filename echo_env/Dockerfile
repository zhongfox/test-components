FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/echo_env

COPY ./ /go/src/echo_env

RUN set -ex && \
go build -v -o /go/bin/echo_env \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/echo_env /usr/bin/
CMD ["echo_env"]
