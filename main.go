package main

import (
	"gineasy/conf"
	"gineasy/routers"
)

func main() {
	r := routers.InitRouter()
	_ = r.Run(conf.AppPort)
}
