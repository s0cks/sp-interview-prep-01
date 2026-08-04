# Interview Prep 01

> This repository contains a small project I have constructed over the course of a weekend to prepare for an upcoming interview.
>
> **Start**: Friday, July 31, 2026
>
> **End**: Monday, August 3, 2026

## Overview

The following project documents the result of prompting Claude for help practicing for
an upcoming interview and then asking it to give me a project to work on as prep.

### Initial Prompt

i have a technical interview (2nd interview) for a role and have been laid off for a while.
can you help me
role is attached.

---

***Couple rounds of a mock interview w/ Claude later***

---

### Practice Project Prompt

give me a go proj for the weekend

### Story

Given exactly what tripped you today — long-running goroutine loops, single-consumer channel discipline, and a real REST API surface (their actual stack) — here's a project that hits all three without ballooning into a weekend-eating monster:

**Build a sensor telemetry service in Go**, close to the exact problem we just live-coded, but complete:

1. **Ingestion:** An HTTP endpoint (`POST /readings`) that accepts a `SensorReading` JSON body and pushes it onto a channel. This gets you real `net/http` reps, which is directly their stack.
2. **Processing:** A pool of per-sensor worker goroutines (spawned on-demand, like you designed today) that maintain a genuine sliding window — actual ring buffer this time, not a tumbling batch — and recompute a rolling average as each reading arrives.
3. **Query:** A `GET /sensors/{id}/average` endpoint that reads the current rolling average — this forces you to solve "how does an HTTP handler safely read state a goroutine is concurrently writing" (mutex-per-sensor, or a request/response channel pattern — pick one and be able to justify it).
4. **Shutdown:** Wire up `context.Context` for clean cancellation instead of the `<-readings` hack from earlier — spin it up with `context.WithCancel` in `main`, pass it through to your workers, handle `os.Interrupt` via `signal.NotifyContext`. This directly fixes the exact gap I flagged today.
5. **Tests:** Since you're TDD-minded, write the ring buffer and averaging logic test-first, before wiring up HTTP — good practice and gives you a concrete "here's tested code" artifact if it comes up.

Stretch, only if you've got extra time: swap the in-memory map for something that mimics their air-gapped/cloud story — e.g., an interface for storage with two implementations (in-memory vs. something pluggable), so you can talk about the config-driven deployment pattern you pitched Tuesday with actual code behind it, not just words.

Scope it to a few hours, not a full system — the goal is muscle memory on the specific patterns, not a finished product. Good luck this weekend.

## Architecture

> You can find some architectural docs here:

See [Architecture](/docs/Architecture.md).

## Building

> You can find how to build the various components here:

See [Building](/docs/Building.md).

## Building All the Docker Images

### Prerequisites

- [task](https://taskfile.dev) -- v3.51.1+ `Optional, useful for building`
- [docker-compose](https://docs.docker.com/compose/) -- v5.1.4+

---

Using **docker-compose**:

```sh
docker compose build
```

Using **task**:

```sh
task build:images
```

---

This builds images for the following:

- The [service](https://github.com/s0cks/sp-interview-prep-01/pkgs/container/sp-interview-prep-01-sensor-service) image.
- The [web](https://github.com/s0cks/sp-interview-prep-01/pkgs/container/sp-interview-prep-01-sensor-web) image.
- The [mock-sensor](https://github.com/s0cks/sp-interview-prep-01/pkgs/container/sp-interview-prep-01-sensor-mock) image.

See: [docker-compose.yaml](/docker-compose.yaml)

## Timeline

> You can find a timeline of working sessions here:

See [Timeline](/docs/Timeline.md).

## LICENSE

See [LICENSE](/LICENSE)

## Credits

- [Claude](https://claude.ai) --- For helping with practice, project idea and feedback.
