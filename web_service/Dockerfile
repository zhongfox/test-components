FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/web_service

COPY ./ /go/src/web_service

RUN set -ex && \
go build -v -o /go/bin/web_service \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/web_service /usr/bin/
EXPOSE 3000
CMD ["web_service"]
