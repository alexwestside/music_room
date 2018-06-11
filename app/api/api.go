package api

import (
	"github.com/kataras/iris"
	"github.com/music_room/serverHTTP"
)


func Routing() {

	serverHTTP.Server.App.Get("/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from Music-Room SERVER")
	})

	serverHTTP.Server.App.PartyFunc("/room", func(r iris.Party) {
		r.Get("/welcome", func(ctx iris.Context) {
			ctx.Writef("Hello from ROOM")
		})
	})


}