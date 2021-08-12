package main

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/application"
	"github.com/vlsidlyarevich/load-balancer/pkg/proxy"
	"log"
)

func main() {
	var serv = application.NewApplication(":8082")

	serv.RouteHandler(proxy.HelloBalancedProxy())

	log.Fatal(serv.Start())
}
