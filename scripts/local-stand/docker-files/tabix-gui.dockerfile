FROM node:18.17.1-slim as builder

WORKDIR /app

RUN apt update && apt -y install git && git clone -b master https://github.com/tabixio/tabix.git /app \
 && echo 'nodeLinker: node-modules' > .yarnrc.yml && yarn set version 3.1.1 \
 && yarn -v && yarn install && yarn build

FROM nginx:1.25-alpine3.18

COPY --from=builder /app/dist /usr/share/nginx/html

# docker run --name tabix-gui-nginx -p 8080:80 -v $PWD:/usr/share/nginx/html:ro nginx:1.25-alpine3.18
