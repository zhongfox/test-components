FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/env_to_output

COPY ./ /go/src/env_to_output

RUN set -ex && \
go build -v -o /go/bin/env_to_output \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/env_to_output /usr/bin/
CMD ["env_to_output"]
