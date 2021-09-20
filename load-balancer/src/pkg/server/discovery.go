package server

import (
	"encoding/json"
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

type Config struct {
	Discovery struct {
		Servers []struct {
			Ip     string `json:"ip"`
			Weight int    `json:"weight"`
		} `toml:"servers" validate:"required"`
	} `toml:"discovery"`
}

type EnvConfig struct {
	Servers string `envconfig:"LOAD_BALANCER_SERVER_LIST"`
}

func NewConfigBasedDiscovery(configPath string) *ConfigBasedDiscovery {
	return &ConfigBasedDiscovery{
		servers: toServers(*load(configPath)),
	}
}

func (d ConfigBasedDiscovery) ServerList() []*Server {
	return d.servers
}

func toServers(c Config) []*Server {
	var result []*Server
	for _, x := range c.Discovery.Servers {
		result = append(result, newServer(x.Ip, x.Weight))
	}

	return result
}

func load(path string) *Config {
	if path == "" {
		log.Panicf("Invalid config path!")
	}
	log.Println("Loading server list from path:", path)

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	if err := envOverride(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

//TODO write custom decoder
func envOverride(c *Config) error {
	var envOverride = new(EnvConfig)
	err := envconfig.Process("LOAD_BALANCER", envOverride)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal([]byte(envOverride.Servers), &c.Discovery.Servers)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}
