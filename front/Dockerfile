FROM node:14.7.0-alpine

WORKDIR /app

RUN apk update && \
    npm install -g npm && \
    npm install -g create-react-app && \
    npm install -g eslint@6.6.0

COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json

RUN npm install

COPY . .

EXPOSE 3000
