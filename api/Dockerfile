FROM golang:1.16.3-alpine3.13

RUN RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirrors.gigenet.com/g' /etc/apk/repositories

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
