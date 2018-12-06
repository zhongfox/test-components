FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/echo_numbers

COPY ./ /go/src/echo_numbers

RUN set -ex && \
go build -v -o /go/bin/echo_numbers \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/echo_numbers /usr/bin/
CMD ["echo_numbers"]

LABEL TencentHubComponent='{\
  "input": [ {"name": "FROM"}, {"name": "TO"}]\
}'
