ARG NODE_VERSION=20-alpine
FROM node:${NODE_VERSION} AS build

ARG SENSOR_SERVICE_URL=http://localhost:8080

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

ENV SENSOR_SERVICE_URL=${SENSOR_SERVICE_URL}
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost/ || exit 1

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
