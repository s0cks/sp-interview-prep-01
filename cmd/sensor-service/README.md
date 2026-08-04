# sensor-service

> A REST service for sensors to publish data to and provide metrics on

## Building

### Prerequisites

- [go](https://go.dev) -- v1.26.4+
- [task](https://taskfile.dev) -- v3.51.1+ `Optional, useful for building`

Using **go**:

```sh
go build \
    -o ./bin/sensor-service \
    ./cmd/sensor-service
```

Using **task**:

```sh
task build:service
```

### Building the Docker Image

Using **compose**:

```sh
docker compose build service
```

## Running Locally

```sh
go run \
    ./cmd/sensor-service
```
