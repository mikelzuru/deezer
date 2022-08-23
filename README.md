# deezer

This repo contains a gRPC server and client implementation to interact with Deezer's API.

## Setup

To setup your system, do the following:

```bash
make setup-dev
```

Should you make any changes to the code, do the following in order to check them:

```bash
make sanity-check
```

## Run

### Server

To interac to Deezer API, first run the server side:

```bash
go run cmd/main.go serve
```

### Client

To interact with the server, use client side command:

```bash
go run cmd/main.go search
```

The above command will search 'hasselhoff' by default. In order to enter your own input, do the following:

```bash
go run cmd/main.go search --query ACDC
```
