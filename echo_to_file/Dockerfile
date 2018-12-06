FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/echo_to_file

COPY ./ /go/src/echo_to_file

RUN set -ex && \
go build -v -o /go/bin/echo_to_file \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/echo_to_file /usr/bin/
CMD ["echo_to_file"]

LABEL TencentHubComponent='{\
  "input": [ {"name": "CONTENT"}, {"name": "FILE"}]\
}'
