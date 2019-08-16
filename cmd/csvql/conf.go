package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Port string
}

var cf Configuration

func loadConfig(fn string) {
	cf.Port = "3306"
	if fn == "" {
		return
	}
	f, err := os.Open(fn)
	assert(err)
	defer f.Close()
	yd := yaml.NewDecoder(f)
	assert(yd.Decode(&cf))
}
