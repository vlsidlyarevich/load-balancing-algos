package main

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/proxy"
	"github.com/vlsidlyarevich/load-balancer/pkg/server"
	"log"
)

func main() {
	var serv = server.NewServer(":8082")

	serv.RouteHandler(proxy.SimpleBalancedProxy())

	log.Fatal(serv.Start())
}
