FROM alpine

MAINTAINER zhongfox (https://github.com/zhongfox)

RUN apk add --no-cache --update --verbose grep bash nmap-ncat && \
  rm -rf /var/cache/apk/* /tmp/* /sbin/halt /sbin/poweroff /sbin/reboot

COPY server.sh /

ENTRYPOINT ["/bin/sh", "/server.sh"]
