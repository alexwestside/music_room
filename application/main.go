package application

import (
	"github.com/kataras/iris"
	"github.com/music_room/serverHTTP"
	"github.com/music_room/application/api"
)

func Run() {

	api.Routing()

	errServer := serverHTTP.Server.App.Run(
		// Start the web serverHTTP at localhost:80
		iris.Addr(":8000"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err serverHTTP closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

	if errServer != nil {
		panic(errServer)
	}

}