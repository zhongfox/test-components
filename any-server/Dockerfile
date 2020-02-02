FROM golang:1.13-alpine as builder

MAINTAINER foxzhong@tencent.com
WORKDIR /go/src/any-server
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on go mod download

COPY . .
RUN set -ex && go build -v -o /go/bin/any-server -gcflags '-N -l' ./*.go

FROM alpine:3.8
COPY --from=builder /go/bin/any-server /usr/bin/
#EXPOSE 3000
CMD ["any-server"]
