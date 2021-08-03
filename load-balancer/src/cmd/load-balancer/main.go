package main

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"log"
)

func main() {
	var serv = server.NewServer(":8081")
	log.Fatal(serv.Start())
}
