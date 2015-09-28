package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type Config struct {
	Server struct {
		Addr string
		Port string
	}
	Service struct {
		Env         string
		Private_Key string
		Issuer      string
	}
}

func LoadConfig() Config {
	cfg := Config{}
	err := gcfg.ReadFileInto(&cfg, "auth_config.gcfg")
	if err != nil {
		fmt.Println("Error while loading config: %s", err.Error())
		return Config{}
	}
	return cfg
}

func privateKey() []byte {
	return []byte(LoadConfig().Service.Private_Key)
}
