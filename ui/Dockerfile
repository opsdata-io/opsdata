FROM node:15.14.0-alpine3.13 as build-stage
WORKDIR /app
COPY package*.json /app/
#RUN npm install
RUN yarn install --production
COPY ./ /app/
RUN npm run build

FROM nginx:1.20.0-alpine
COPY --from=build-stage /app/build/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
