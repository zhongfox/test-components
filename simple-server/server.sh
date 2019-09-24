#!/bin/sh

tcp_server() {
  echo "start tcp server on port $1"
  ncat -k -l -p $1 -c 'echo -e "tcp: ${TCP_CONTENT}"';
}

http_server() {
  echo "start http server on port $1"
  ncat -k -l -p $1 -c 'echo -e "HTTP/1.1 200 OK\n\n http: ${HTTP_CONTENT}"';
}

if [[ ! -z "${TCP_PORT}" ]] && [[ ! -z "${HTTP_PORT}" ]]; then
  trap 'kill -TERM $BGPID; exit' SIGINT
  tcp_server ${TCP_PORT} &
  BGPID=$!
  http_server ${HTTP_PORT}
elif  [[ ! -z "${TCP_PORT}" ]]; then
  tcp_server ${TCP_PORT}
elif  [[ ! -z "${HTTP_PORT}" ]]; then
  http_server ${HTTP_PORT}
else
  echo "no server started, env TCP_PORT or HTTP_PORT required."
fi
