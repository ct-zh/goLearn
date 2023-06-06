package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 通用config配置文件读取方法

type Config struct {
	Path string `yaml:"path"`
} // 通用配置

func NewConfig() *Config {
	cfg := &Config{}
	cfgFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
