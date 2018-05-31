package main

import (
	"github.com/kataras/iris"
	"github.com/music_room/adminpanel"
)

func main() {

	go adminpanel.Run()

	app := iris.New()
	app.Logger().SetLevel("debug")


	app.PartyFunc("/v1", func(r iris.Party) {
		r.Get("/welcome", func(ctx iris.Context) {
			ctx.Writef("Hello from Music-Room SERVER")
		})
	})

	app.Run(
		// Start the web server at localhost:80
		iris.Addr(":8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)


	//config.Load()
	//
	//server := gin.Default()
	//
	//mux := adminpanel.New()
	//
	//server.GET("/adminpanel/", gin.WrapH(mux))
	//
	//v1 := server.Group("/v1")
	//{
	//	v1.GET("/welcome", func(c *gin.Context) {
	//		firstname := c.DefaultQuery("firstname", "Guest")
	//		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//
	//		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//	})
	//}
	//
	//server.Run("localhost:8888")

}
