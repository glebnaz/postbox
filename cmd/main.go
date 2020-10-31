package main

import (
	"fmt"

	"github.com/glebnaz/postbox/cmd/config"
	"github.com/glebnaz/postbox/cmd/server"
)

const banner = `
Hi, i'm your postman!
 _______
/O     O\
  \ - /

Created by GlebNaz!

Version: %s
Build: %s
`

var (
	//Version is lbsflag
	Version = "NO VERSION SET"
	//Build is lbsflag
	Build = "NO BUILD SET"
)

func main() {
	fmt.Printf(banner, Version, Build)
	var conf config.Config
	conf.Init()

	app := server.InitServer(conf.DBURL, conf.User, conf.Pass)
	app.Run(conf.PORT)
}
