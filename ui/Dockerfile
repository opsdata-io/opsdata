FROM node:15.14.0-alpine3.13
RUN mkdir -p /opt/app
WORKDIR /opt/app
COPY . .
RUN npm run build
RUN npm install -g serve
EXPOSE 5000
CMD [ "serve", "-s", "build" ]
