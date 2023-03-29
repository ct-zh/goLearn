package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	OpenKey string `json:"open_key"`
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
	if cfg.OpenKey == "" {
		panic("open_key is empty")
	}

	return cfg
}
