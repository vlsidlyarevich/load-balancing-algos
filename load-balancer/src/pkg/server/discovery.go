package server

import (
	"github.com/BurntSushi/toml"
	"log"
)

const DefaultConfigPath = "conf/config.toml"

type Discovery interface {
	ServerList() []*Server
}

type ConfigBasedDiscovery struct {
	servers []*Server
}

type config struct {
	servers []string
}

func NewConfigBasedDiscovery() *ConfigBasedDiscovery {
	return &ConfigBasedDiscovery{
		servers: toServers(load(DefaultConfigPath)),
	}
}

func toServers(urls []string) []*Server {
	var result []*Server
	for _, x := range urls {
		result = append(result, newServer(x))
	}

	return result
}

func load(path string) []string {
	if path == "" {
		log.Panicf("Invalid config path!")
	}
	log.Println("Loading server list from path:", path)

	var config config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	return config.servers
}

func (d ConfigBasedDiscovery) ServerList() []*Server {
	return d.servers
}
