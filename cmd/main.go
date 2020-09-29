package main

import (
	"github.com/glebnaz/postbox/cmd/config"
	"github.com/glebnaz/postbox/cmd/server"
)

var (
	Version = "NO VERSION SET"
	Commit  = "NO COMMIT SET"
	Build   = "NO BUILD SET"
)

func main() {
	var conf config.Config
	conf.Init()

	server.InitServer(conf.DBURL, conf.User, conf.Pass)
}
