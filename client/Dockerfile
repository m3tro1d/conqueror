FROM node:16.17.1-alpine

ENV PATH /app/node_modules/.bin:$PATH

WORKDIR /app

COPY package.json .
COPY yarn.lock .
RUN yarn install --frozen-lockfile

RUN mkdir node_modules/.cache && chmod -R 777 node_modules/.cache

COPY . .

EXPOSE 3000

CMD [ "npm", "start" ]
