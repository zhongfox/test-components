FROM golang:1.10-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/job_out

COPY ./ /go/src/job_out

RUN set -ex && \
go build -v -o /go/bin/job_out \
-gcflags '-N -l' \
./*.go

FROM alpine:3.8
RUN mkdir -p /root/src
WORKDIR /root/src
COPY --from=builder /go/bin/job_out /usr/bin/
CMD ["job_out"]

LABEL TencentHubComponent='{\
  "description": "简易输出",\
  "input": [ {"name": "OUTPUTS", "desc": "可选, 需要输出的内容, 多个用换行分隔, kv 用=分隔"} ],\
  "output": []\
}'
