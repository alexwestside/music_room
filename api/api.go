package api

import (
	"github.com/music_room/api/v1"
	"github.com/kataras/iris"
	"github.com/music_room/api/v2"
)

func Routing(app *iris.Application) {

	app.PartyFunc("/v1", func(r iris.Party) {
		v1.Routes(r)
	})

	app.PartyFunc("/v2", func(r iris.Party) {
		v2.Routes(r)
	})
}