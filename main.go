package main

import (
	"OnlinePhotoAlbum/conf"
	"OnlinePhotoAlbum/routers"
)

func main() {
	r := routers.InitRouter()
	_ = r.Run(conf.Port)
}
