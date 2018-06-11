package main

import (
	"github.com/music_room/serverHTTP"
	"github.com/music_room/app"
	"github.com/music_room/adm"
)


func main() {

	serverHTTP.New()

	adm.Run()

	app.Run()

}
