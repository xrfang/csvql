package main

import (
	"flag"
	"gopkg.in/src-d/go-mysql-server.v0"
	"gopkg.in/src-d/go-mysql-server.v0/auth"
	"gopkg.in/src-d/go-mysql-server.v0/server"
)

func main() {
	conf := flag.String("conf", "", "configuration file")
	flag.Parse()
	loadConfig(*conf)
	e := sqle.NewDefault()
	assert(e.Init())
	an := auth.NewNativeSingle(cf.User, cf.Pass, auth.ReadPerm)
	cfg := server.Config{
		Protocol: "tcp",
		Address:  "0.0.0.0:" + cf.Port,
		Auth:     an,
	}
	s, err := server.NewDefaultServer(cfg, e)
	assert(err)
	assert(s.Start())
}
