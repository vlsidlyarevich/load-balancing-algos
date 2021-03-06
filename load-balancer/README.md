## Echo server
Simple go-written echo-server

## Development plan
1) Simple load balancer server with hardcoded servers - DONE
2) Implement simple "Round robin" algo - DONE
3) Fetch server list from config - DONE
4) Implement "Weighted Round robin" algo
5) Implement algo switch functionality
6) Implement server health checks with goroutines

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to install Golang and Docker to run this project.

Golang:
``` bash 
✗ go version
go version go1.15.6 darwin/amd64
```

Docker:
``` bash
✗ docker -v
Docker version 20.10.7, build f0df350
```

### Installation

* Clone the project and open in your favourite IDE.
* Install all needed dependencies via `go mod download` or via IDE tools.
* Build Docker image by running `docker build . -t load-balancer`


## Running

* Run [main.go](src/cmd/echo-server/main.go)
* Run containerized `docker run -p 8080:8080 load-balancer` where `8080` - default port
