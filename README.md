# Hathor
Go service developed following clean architecture. Centered around the concept of events, provides a GRPC layer to perform CRUD operations on them.

For a deeper look into the service architecture, including app structure, design decisions, library choices and more take look at [ARCHITECTURE.md](ARCHITECTURE.md).

## How to run
Docker and docker-compose are needed to start up the server.

`make up` will build the containers from scratch, starting the DB and, if successful, the server on port 3333.

A `.env` file is needed, provided upon request. That said, for testing purposes, something basic like this would work:

```
SERVER_HOST=localhost
SERVER_PORT=2222
DB_NAME=hathor
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_HOST=postgres
DB_PORT=5432
```

## Test the server
Some basic curl commands are provided for convenience to actually test the server is operating right.

### PING
```bash
grpcurl -plaintext \
    localhost:2222 \
    hathor.Hathor.Ping
```

### POST
```bash
grpcurl -plaintext \
    -d '{"name": "Random Event", "description": "a random test event", "category": "random", "location": "Buenos Aires", "publisher": "Nico", "lineup": ["Random Artits"], "start_ts": "2023-01-01T18:00:00Z","end_ts": "2023-01-01T21:00:00Z"}' \
    localhost:2222 \
    hathor.Hathor.CreateEvent
```

### GET
```bash
grpcurl -plaintext \
    -d '{"event_id": 1}' \
    localhost:2222 \
    hathor.Hathor.GetEvent
```

### UPDATE
```bash
grpcurl -plaintext \
    -d '{"update_mask": {"paths": "event.lineup"}, "event": {"id": 1, "lineup": ["Random", "Artists", "Changed"]}}' \
    localhost:2222 \
    hathor.Hathor.UpdateEvent
```

### DELETE
```bash
grpcurl -plaintext \
    -d '{"event_id": 1}' \
    localhost:2222 \
    hathor.Hathor.DeleteEvent
```

## TBD
1. Deployment github workflow.
