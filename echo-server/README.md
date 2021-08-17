## Echo server
Simple go-written echo-server

## Development plan
1) Simple echo server
2) Fetching name and port from config
3) Implement /health endpoint
TBD

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

### Installing

* Clone the project and open in your favourite IDE.
* Install all needed dependencies via `go mod download` or via IDE tools.
* Build Docker image by running `docker build . -t echo-server`


## Running

* Run [main.go](src/cmd/echo-server/main.go)
* Run containerized `docker run -p 8081:8081 echo-server`
