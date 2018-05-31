package music_room_server

import "github.com/kataras/iris"


func New() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

func Run(app *iris.Application) {
	app.Run(
		// Start the web music_room_server at localhost:80
		iris.Addr(":8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err music_room_server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}