package server

import (
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Discovery interface {
	ServerList() []*Server
}

type ConfigBasedDiscovery struct {
	servers []*Server
}

type config struct {
	discovery discoveryConf
}

type discoveryConf struct {
	servers []string `validate:"required" envconfig:"LOAD_BALANCER_SERVER_LIST"`
}

func NewConfigBasedDiscovery(configPath string) *ConfigBasedDiscovery {
	return &ConfigBasedDiscovery{
		servers: toServers(*load(configPath)),
	}
}

func (d ConfigBasedDiscovery) ServerList() []*Server {
	return d.servers
}

func toServers(c config) []*Server {
	var result []*Server
	for _, x := range c.discovery.servers {
		result = append(result, newServer(x))
	}

	return result
}

func load(path string) *config {
	if path == "" {
		log.Panicf("Invalid config path!")
	}
	log.Println("Loading server list from path:", path)

	var config config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	if err := envOverride(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func envOverride(c *config) error {
	err := envconfig.Process("LOAD_BALANCER", c)
	if err != nil {
		return err
	}

	return nil
}
