package main

import (
	"fmt"
	"github.com/glebnaz/postbox/cmd/config"
	"github.com/glebnaz/postbox/cmd/server"
)

const banner = `
Hi, i'm your postman!

Version: %s
Build: %s
____________________________________O/_______
                                    O\
`

var (
	Version = "NO VERSION SET"
	Commit  = "NO COMMIT SET"
	Build   = "NO BUILD SET"
)

func main() {
	fmt.Printf(banner, Version, Build)
	var conf config.Config
	conf.Init()

	app := server.InitServer(conf.DBURL, conf.User, conf.Pass)
	app.Run(conf.PORT)
}
