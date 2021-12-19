# Project Name

Project summary

## Input output detail

TODO

## Code structure

### `cmd`

Main executable: `cmd/example`

### `conf`

Environment variables for initializing the app.

### `pkg/core`
Business logic. Can be tested without external resources (database,
HTTP, websocket, message queue, file, ..)

### `pkg/driver`
Calls to external resources.
