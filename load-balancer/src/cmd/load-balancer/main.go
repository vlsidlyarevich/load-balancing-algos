package main

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/proxy"
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"log"
)

func main() {
	var serv = server.NewServer(":8081")

	serv.AddProxy(proxy.NewHelloProxy())

	log.Fatal(serv.Start())
}
