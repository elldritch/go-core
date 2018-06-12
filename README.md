# go-core

A personal collection of useful application-level packages for writing HTTP API
servers.

My usual stack:

- go-chi/chi: mux
- rs/zerolog: structured logging
- volatiletech/sqlboiler: "ORM"
- jmoiron/sqlx: database driver sugar
- lib/pq: database driver
- contribsys/faktory: task queue

Things I still need:

- Distributed tracing
- Metrics and statistics
- RPC
- Dependency injection
