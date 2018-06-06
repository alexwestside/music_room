package api

import (
	"github.com/music_room/application/api/v1"
	"github.com/kataras/iris"
	"github.com/music_room/application/api/v2"
	"github.com/music_room/serverHTTP"
)

func Routing() {

	serverHTTP.Server.App.PartyFunc("/v1", func(r iris.Party) {
		v1.Routes(r)
	})

	serverHTTP.Server.App.PartyFunc("/v2", func(r iris.Party) {
		v2.Routes(r)
	})
}