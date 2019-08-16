package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	User string
	Pass string
	Port string
}

var cf Configuration

func loadConfig(fn string) {
	cf.User = "dev"
	cf.Pass = "Password01!"
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
