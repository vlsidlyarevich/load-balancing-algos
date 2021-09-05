package main

import (
	"github.com/vlsidlyarevich/load-balancer/pkg/application"
	"github.com/vlsidlyarevich/load-balancer/pkg/proxy"
	"log"
	"os"
)

const DefaultConfigPath = "conf/config.toml"

func main() {
	var args = os.Args[1:]
	log.Println("Arguments used:", args)

	var configPath = DefaultConfigPath
	if len(args) > 0 && args[0] != "" {
		configPath = args[0]
	}

	var serv = application.NewApplication(":8080", args, configPath)

	//TODO clojure?
	serv.RouteHandler(proxy.HelloBalancedProxy(serv.Env))

	log.Fatal(serv.Start())
}
