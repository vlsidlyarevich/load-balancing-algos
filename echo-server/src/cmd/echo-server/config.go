package main

import (
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	server ServerConf
}

type ServerConf struct {
	port string `validate:"required" envconfig:"ECHO_SERVER_PORT"`
	id   string `validate:"required" envconfig:"ECHO_SERVER_ID"`
}

func Load(path string) *Config {
	var config Config
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatal(err)
	}

	if err := envOverride(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func envOverride(c *Config) error {
	err := envconfig.Process("echo-server", c)
	if err != nil {
		return err
	}

	return nil
}
