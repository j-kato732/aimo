FROM node:14.0.0-alpine3.11 as build-stage
WORKDIR /app
COPY front/app/aimo/package*.json ./
RUN npm install
COPY front/app/aimo/ ./
RUN npm run build

FROM nginx:1.17.10-alpine as production-stage
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY nginx/nginx.conf /etc/nginx/nginx.conf