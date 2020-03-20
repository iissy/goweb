package main

import (
	_ "hrefs.cn/src"
	"hrefs.cn/src/api"
	"hrefs.cn/src/srv"
	"hrefs.cn/src/web"
)

func main() {
	go api.Start()
	go web.Start()
	srv.Start()
}
