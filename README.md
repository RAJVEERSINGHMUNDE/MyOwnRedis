# MyOwnRedis

A small Redis-inspired server written in Go. This is a personal learning project and mostly just me messing around with TCP servers, the RESP wire protocol, in-memory data structures, and append-only persistence.

The reference for this was a nice blog I found online https://www.build-redis-from-scratch.dev/

## What it does

- Listens for TCP connections on port `6379`.
- Parses RESP arrays and bulk strings.
- Supports a small command set: `PING`, `SET`, `GET`, `HSET`, and `HGET`.
- Stores string keys and hash fields in memory.
- Writes `SET` and `HSET` requests to `database.aof` and syncs the file roughly once per second.

## Running it

You need Go 1.26.5 or a compatible Go installation.

```bash
cd src
go run .
```

The server starts on standard redis port 6379. You can connect using the actual redis-cli, you might face problems with request routing(assuming you are using WSL) since redis-cli comes prebuilt only for linux. 

```bash
redis-cli -p 6379 PING
redis-cli -p 6379 SET greeting hello
redis-cli -p 6379 GET greeting
redis-cli -p 6379 HSET user:1 name Ada
redis-cli -p 6379 HGET user:1 name
```

## Supported commands

| Command | Arguments | Result |
| --- | --- | --- |
| `PING [message]` | Optional message | `PONG`, or the supplied message |
| `SET key value` | Key and value | Stores a string and returns `OK` |
| `GET key` | Key | Returns the stored string or a null bulk string |
| `HSET hash field value` | Hash name, field, value | Stores a hash field and returns `OK` |
| `HGET hash field` | Hash name and field | Returns the stored value or a null bulk string |

## Scope and limitations

- accepts one client connection at a time;
- has no authentication, configuration, expiration, transactions, pub/sub, replication, clustering, or eviction;
- only implements the RESP value types needed for its current commands;
- keeps all live data in memory;
- records writes to an AOF file but does not yet replay that file when starting up;
- does not aim for Redis protocol or command compatibility beyond the small surface above.


## Project layout

```text
src/
  main.go       TCP server and command dispatch
  resp.go       RESP parsing and response encoding
  handler.go    Command handlers and in-memory storage
  aof.go        Append-only-file writing and periodic sync
```
Other files are irrelevant.

