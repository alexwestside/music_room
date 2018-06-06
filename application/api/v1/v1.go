package v1

import (
	"github.com/kataras/iris"
)

func Routes(r iris.Party) {
	r.Get("/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from Music-Room SERVER")
	})
}
