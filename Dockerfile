FROM golang:1.16.3-alpine3.13

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
