package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type yml1 struct {
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"pwd"`
	}
}

func main() {
	filePtr, err := os.Open("./conf2.yml")
	if err != nil {
		panic(err)
	}
	defer filePtr.Close()

	var res yml1

	decoder := yaml.NewDecoder(filePtr)
	err = decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", res)
}
