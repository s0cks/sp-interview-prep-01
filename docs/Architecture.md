# Architecture

> The following flowchart briefly describes the architecture of the service.

```mermaid
%%{init: { 'flowchart': { 'padding': 40, 'nodeSpacing': 50, 'rankSpacing': 60 } } }%%
flowchart LR
    WEB["Web"]
    CLIENT["Client"]
    SENSOR["Sensor"]

    SERVICE_IN(["Service"])

    subgraph BUFFER_REG_INMEM["Buffer Registry (In-Memory)"]
        direction TD
        SB0[("Sensor Buffer #1")]
        SB1[("Sensor Buffer #2")]
        SBN[("Sensor Buffer #N")]
    end

    SERVICE_IN <--> BUFFER_REG_INMEM

    CLIENT & WEB <--> | GET /sensors/:id/metric | SERVICE_IN
    SENSOR <--> | POST /sensors/:id | SERVICE_IN
```

- A client or sensor will communicate over REST to the Go service.

- The service itself will buffer data POST by a sensor in a ring buffer.

  - The ring buffer capacity specified by a process argument, but defaults to 10.
  - I chose a ring buffer for its efficient use of space in time series like data where you only need the last N items.
  - buffers are only created once a sensor POSTs data to the service ensuring that there are never any empty buffers.

- The service will use a read-write mutex per ring buffer, this is because:

  - A read-write mutex will allow us to support many concurrent reads, and only blocking exclusively when writing data.
  - A read-write mutex per buffer allows each sensor and buffer to block independently of each other.

## Endpoints

I have created sequence diagrams so I can plan and document how the various
endpoints will work:

### POST /sensors/{id}

> A sequence diagram for how the POST endpoint works

Sensors can post data for the service to ingest here

```mermaid
sequenceDiagram
    autonumber
    participant Sensor as Sensor
    participant Service as Service
    participant Registry as Sensor Buffer Registry
    participant Buffer as Sensor Buffer

    Sensor->>Service: POST /sensors/{id}
    Service->>Registry: Get buffer for {id}
    Note over Registry: sync.Map.LoadOrStore()
    Registry->>Service: buffer

    Service->>Buffer:
    Note over Buffer: rwlock.Lock()
    Note over Buffer: write data
    Note over Buffer: rwlock.Unlock()
    Buffer->>Service: 
    Service->>Sensor: 200
```

#### GET /sensors

> A Sequence diagram for how the GET sensor endpoint works

Returns a list of sensor ids

```mermaid
sequenceDiagram
    autonumber
    participant Client as Client
    participant Service as Service
    participant Registry as Sensor Buffer Registry
    participant Buffer as Sensor Buffer

    Client->>Service: GET /sensors
    alt Buffer registry is empty
        Registry->>Service: nil
        Service->>Client: 204
    else
        Registry->>Service: buffer
    end

    Service->>Buffer:
    Note over Buffer: rwlock.RLock()
    Note over Buffer: read data as copy
    Note over Buffer: rwlock.RUnlock()
    Buffer->>Service: data
    Service->>Client: 200
```

#### GET /sensors/{id}

> A Sequence diagram for how the GET sensor endpoint works

This endpoint returns all the buffered data for the sensor

```mermaid
sequenceDiagram
    autonumber
    participant Client as Client
    participant Service as Service
    participant Registry as Sensor Buffer Registry
    participant Buffer as Sensor Buffer

    Client->>Service: GET /sensors/{id}

    Service->>Registry: Get buffer for {id}
    Note over Registry: sync.Map.Load()
    alt Buffer {id} not found
        Registry->>Service: nil
        Service->>Client: 204
    else
        Registry->>Service: buffer
    end

    Service->>Buffer:
    Note over Buffer: rwlock.RLock()
    Note over Buffer: read data as copy
    Note over Buffer: rwlock.RUnlock()
    Buffer->>Service: data
    Service->>Client: 200
```

#### GET /sensors/{id}/{metric}

> A sequence diagram for how the GET metric endpoint works

This endpoint returns the calculated value for the corresponding metric.

It supports the following metrics:

- average
- sum

We calculate the result on demand to improve ingestion time and provide flexibility and scaling for any future metrics.

```mermaid
sequenceDiagram
    autonumber
    participant Client as Client
    participant Service as Service
    participant Registry as Sensor Buffer Registry
    participant Buffer as Sensor Buffer

    Client->>Service: GET /sensors/{id}/{metric}
    alt Invalid metric
        Service->>Client: 400
    end

    Service->>Registry: Get buffer for {id}
    Note over Registry: sync.Map.Load()
    alt Buffer not found
        Registry->>Service: nil
        Service->>Client: 204
    end

    Service->>Buffer:
    Note over Buffer: rwlock.RLock()
    Note over Buffer: Calculate {metric}
    Note over Buffer: rwlock.RUnlock()
    Buffer->>Service: results
    Service->>Client: 200
```
