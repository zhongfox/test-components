FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/wait_seconds

COPY ./ /go/src/wait_seconds

RUN set -ex && \
go build -v -o /go/bin/wait_seconds \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/wait_seconds /usr/bin/
CMD ["wait_seconds"]

LABEL TencentHubComponent='{\
  "description": "简易暂停",\
  "input": [ {"name": "SECOND", "desc": "可选, 暂停秒数, 默认60秒"} ],\
  "output": []\
}'
