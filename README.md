# Consistent Hashing Load Balancer

A production-inspired HTTP load balancer built in **Go** implementing **consistent hashing**, **virtual nodes**, and **reverse proxying** for sticky request routing.

The project demonstrates how requests can be deterministically routed to backend servers while minimizing key remapping and maintaining an even distribution across the cluster.

---

## Features

* Consistent Hash Ring
* MurmurHash3 hashing algorithm
* Configurable Virtual Nodes
* Binary Search based O(log n) backend lookup
* Sticky request routing using `X-User-ID`
* Reverse Proxy built using Go's standard library
* Config-driven backend discovery
* Concurrent benchmark tool
* Request distribution and latency analysis

---

## Architecture

```
                    Client
                       │
                       ▼
                HTTP Load Balancer
                       │
             Extract X-User-ID Header
                       │
                       ▼
             MurmurHash3(User ID)
                       │
                       ▼
             Consistent Hash Ring
                       │
             Binary Search Lookup
                       │
                       ▼
              Physical Backend Server
                       │
                       ▼
                Reverse Proxy
                       │
                       ▼
                    Response
```

---

## Why Consistent Hashing?

Traditional load balancing algorithms such as Round Robin distribute requests evenly but do not preserve request affinity.

Consistent hashing ensures that the same routing key (in this project, the `X-User-ID` header) consistently maps to the same backend server while minimizing remapping when servers are added or removed.

---

## Virtual Nodes

Each physical server is represented by multiple virtual nodes on the hash ring.

Benefits:

* Better distribution across the hash space
* Reduced hotspots
* Improved load balancing
* Easier scaling

Current configuration:

* Physical Servers: **3**
* Virtual Nodes per Server: **500**
* Total Ring Positions: **1500**

---

## Routing Flow

1. Client sends an HTTP request.
2. Middleware extracts the `X-User-ID` header.
3. MurmurHash3 hashes the user ID.
4. Binary search locates the nearest clockwise virtual node.
5. Virtual node maps to a physical backend.
6. Reverse proxy forwards the request.
7. Backend response is streamed back to the client.

---

## Configuration

Backends are configured through `config.yaml`.

Example:

```yaml
server:
  port: 8000

backends:
  - id: server1
    host: localhost
    port: 8001

  - id: server2
    host: localhost
    port: 8002

  - id: server3
    host: localhost
    port: 8003

virtual_nodes:
  total: 500
```

---

## Benchmark

### Load Generator

The project includes a concurrent benchmark tool.

Configuration:

* Total Requests: **100,000**
* Concurrent Workers: **150**
* Random User IDs
* Persistent HTTP clients
* Latency measurement
* Backend distribution analysis

### Results

```
Total Requests : 100000
Successful     : 99834
Failed         : 166

Distribution
------------
server1 : 35433
server2 : 33611
server3 : 30790

Latency
-------
Average : 80.94 ms
Median  : 43.00 ms
P95     : 177.92 ms
P99     : 359.31 ms
Max     : 10404.82 ms

Duration
--------
Total Time : 54.2 s
Requests/s : 1846
```

---

## Engineering Decisions

### MurmurHash3

The project originally used FNV-1a for hashing.

Benchmarking revealed uneven request distribution across the ring. After switching to MurmurHash3, request distribution became significantly more balanced.

This decision was based on experimental benchmarking rather than assumption.

---

### Binary Search

Backend lookup uses binary search over the sorted hash ring.

Time Complexity:

* Lookup: **O(log n)**
* Ring Construction: **O(n log n)**

---

### Reverse Proxy

A reusable reverse proxy forwards requests to the selected backend while preserving the original HTTP request.

A shared HTTP transport is used for connection reuse.

---

## Tech Stack

* Go
* net/http
* httputil.ReverseProxy
* MurmurHash3
* YAML Configuration
* Zerolog


---

## Learning Outcomes

This project explores several backend engineering concepts including:

* Distributed systems fundamentals
* Consistent hashing
* Virtual nodes
* Reverse proxying
* HTTP internals
* Concurrent benchmarking
* Performance analysis
* Config-driven application design
* Load balancing algorithms

---

## License

MIT
