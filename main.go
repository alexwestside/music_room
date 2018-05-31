package main

import (
	"github.com/music_room/adminpanel"
	"github.com/music_room/music_room_server"
	"github.com/music_room/api"
)

func main() {

	go adminpanel.Run()

	MusicRoom := music_room_server.New()

	api.Routing(MusicRoom)

	music_room_server.Run(MusicRoom)
}
