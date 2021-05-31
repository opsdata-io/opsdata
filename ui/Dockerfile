FROM node:current-alpine as build-stage
WORKDIR /app
COPY package*.json /app/
RUN yarn install --production
COPY ./ /app/
RUN npm run build

FROM nginx:alpine
RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/http\:\/\/alpine-mirror.support.tools./g' /etc/apk/repositories
RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/alpine-mirror.support.tools./g' /etc/apk/repositories
RUN apk update && apk add --no-cache \
  curl \
  libjpeg-turbo \
  && rm -rf /var/cache/apk/*
COPY --from=build-stage /app/build/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
