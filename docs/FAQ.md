# FAQ

> A brief discussion on Other Possibilities & Lessons Learned

## What Would I Do Differently?

What would I do differently if:

- I had more time..
- Could redo the entire project..

**Some generalized notes**:

- I would probably split this repository into the different components.

### Go Service

- Move the `{id}` path param from `POST /sensors/{id}` to a request header allowing to simplify the api path to `POST /data`
- Make the `{id}` and `{metric}` path params a query string in `GET /sensors/{id}` or `GET /sensors/{id}/{metric}` like `GET /sensors?id={id}&metric=sum` for a more flexible api.
- Cache the number of samples environment variable at boot in the [main.go](/cmd/sensor-service/main.go) for the [Go service](/internal/sensor/service.go) to use
- Add unit tests for the controller code [service.go](/internal/sensor/service.go).

### Web Dashboard

- Better UX
- Probably would have used typescript.
- Add unit tests for [dashboard components](/client/src/components/).

### Mock Sensor

- Support parallel runs of the containers using [docker-compose](/docker-compose.yaml).
- Add more testing patterns to the [mock-sensor](/test/sensor-mock/) and better arguments for customizing the runs.

## Lessons Learned?

### Go Service

- REST in Go, not as fun as I remember. Do community libraries exist to solve this?

### Web Dashboard

- You apparently don't need [redux](https://redux.js.org/) when dealing with [react](https://react.dev/).

### Mock Sensor

- [nimesis](https://mimesis.name/master/), pretty neat. Never used it.
