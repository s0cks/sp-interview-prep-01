# Sensor Dashboard

A simple [react](https://react.dev/) dashboard implementation.

## Building

### Prerequisites

You will need the following to build the dashboard:

- [node](https://nodejs.org/) --- v26.4.0+
- The following package dependencies:
  - [chart.js](https://www.chartjs.org/) --- v4.5.1+
  - [react](https://react.dev/) --- v19.0.0+
  - [vite](https://vite.dev/) --- v6.0.0+
- [pnpm](https://pnpm.io/) --- v10.28.1+ `Optional, preferrable over npm`
- [task](https://taskfile.dev) --- v3.51.1+ `Optional, useful for building`

---

Using **pnpm**:

```sh
pnpm install
pnpm build
```

Using **npm**:

```sh
npm install
npm run build
```

Using **task**:

```sh
task build:web
```

### Building the Docker Image

Using **compose**:

```sh
docker compose build web
```

## Running Locally

You can run the [vite](https://vite.dev/) dev server:

Using **pnpm**:

```sh
pnpm dev
```

Using **npm**:

```sh
npm run dev
```
