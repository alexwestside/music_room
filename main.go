package main

import (
	"github.com/music_room/serverHTTP"
	"github.com/music_room/adm"
	"github.com/music_room/app"
)


func main() {

	serverHTTP.New()

	go adm.Run()

	app.Run()

}
