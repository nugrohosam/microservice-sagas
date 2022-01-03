FROM node:14-alpine

WORKDIR /app
COPY . .

RUN rm .env
RUN mv .env.docker .env
RUN npm install