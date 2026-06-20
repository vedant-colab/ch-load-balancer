# Consistent Hashing Load Balancer

A lightweight HTTP load balancer written in Go that routes requests using consistent hashing.

## Overview

This project explores how distributed systems use consistent hashing to route requests while minimizing key redistribution when servers are added or removed.

The load balancer maintains a hash ring of backend servers and routes requests to the nearest server in the clockwise direction of the ring.

## Features

* Consistent hashing based request routing
* Configurable backend servers
* Fast server lookup using binary search
* Lightweight implementation using Go's standard library

## Getting Started

### Clone the repository

```bash
git clone <repo-url>
cd consistent-hash-lb
```

### Create configuration

```yaml
server:
  port: 8080
  environment: development
```

Save as `config.yaml`.

### Run

```bash
go run cmd/main.go
```

## Project Structure

```text
cmd/
internal/
├── config/
├── hash/
├── ring/
└── server/
```

## How Consistent Hashing Works

1. Backend servers are placed on a hash ring.
2. Request keys are hashed into the same space.
3. The request is routed to the first server encountered while moving clockwise around the ring.
4. Adding or removing a server affects only a subset of keys.

## Tech Stack

* Go
* net/http
* Zerolog
* YAML Configuration
