package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RequestKey string `yaml:"request_key"`
	RequestIP  string `yaml:"request_ip"`
}

func NewConfig() *Config {
	cfg := &Config{}
	cfgFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		panic(err)
	}
	if cfg.RequestKey == "" {
		panic("request_key is empty")
	}
	if cfg.RequestIP == "" {
		panic("request_ip is empty")
	}

	return cfg
}
