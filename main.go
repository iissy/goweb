package main

import (
	_ "github.com/iissy/goweb/src"
	"github.com/iissy/goweb/src/api"
	"github.com/iissy/goweb/src/srv"
	"github.com/iissy/goweb/src/web"
)

func main() {
	go api.Start()
	go web.Start()
	srv.Start()
}
