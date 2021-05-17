FROM golang:1.16.3-alpine3.13

RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/a1-alpine.support.tools/g' /etc/apk/repositories

RUN apk add -q --update && \
  apk add -q \
  bash \
  git \
  curl \
  nodejs \
  npm \
  && rm -rf /var/cache/apk/*

COPY main /
CMD ["/main"]
