package main

import (
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Server ServerConf
}

type ServerConf struct {
	Port string `validate:"required" envconfig:"ECHO_SERVER_PORT"`
	Id   string `validate:"required" envconfig:"ECHO_SERVER_ID"`
}

func LoadConfig(path string) *Config {
	if path == "" {
		log.Panicf("Invalid config path!")
	}
	log.Println("Loading config from path:", path)

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	if err := envOverride(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func envOverride(c *Config) error {
	err := envconfig.Process("ECHO_SERVER", c)
	if err != nil {
		return err
	}

	return nil
}
