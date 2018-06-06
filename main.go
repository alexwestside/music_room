package main

import (
	"github.com/music_room/serverHTTP"
	"github.com/music_room/adminpanel"
	"github.com/music_room/application"
)


func main() {

	serverHTTP.New()

	go adminpanel.Run()

	application.Run()
}
