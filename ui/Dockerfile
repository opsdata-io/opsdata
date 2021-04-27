FROM node:15.14.0-alpine3.13

RUN apk add --update bash && rm -rf /var/cache/apk/*

# Create app directory
WORKDIR /usr/src/app

COPY package*.json ./

RUN npm ci --only=production

COPY . .

EXPOSE 3000
CMD [ "npm", "start" ]