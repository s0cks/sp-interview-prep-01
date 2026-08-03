# Building

## Prerequisites

- [go](https://go.dev) -- v1.26.4
- [node](https://nodejs.org) -- v26.4.0
- [docker-compose](https://docs.docker.com/compose/) -- v5.1.4
- [task](https://taskfile.dev) -- Optional, useful for building

## Building the Service

```sh
go build cmd/sensor-service

# or using task
task build:service
```

## Building the Dashboard

```sh
cd client/

# build using pnpm:
pnpm build

# build using npm:
npm run build

# or using task
task build:web
```

## Building the Docker Images

You can build all the Docker images using docker-compose

```sh
docker compose build

# or using task
task build:images
```

This builds images for the following:

- The service
- The web dashboard
- The mock sensor

See: [docker-compose.yaml](/docker-compose.yaml)
