FROM node:current-alpine as build-stage
WORKDIR /app
COPY package*.json /app/
#RUN npm install
RUN yarn install --production
COPY ./ /app/
RUN npm run build

FROM nginx:alpine
RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories
RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories
RUN apk add --no-cache --update \
  curl \
  libjpeg-turbo \
  && rm -rf /var/cache/apk/*
COPY --from=build-stage /app/build/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
