package music_room_server

import (
	"github.com/kataras/iris"
	"fmt"
	"github.com/jinzhu/gorm"
)

type HTTPHandler struct {
	App     *iris.Application
	Address string
	//db      *gorm.DB
}


func New(host string, port string, connection *gorm.DB) *HTTPHandler {
	return &HTTPHandler {
		Address: fmt.Sprintf("%s:%s", host, port),
		App:     iris.New(),
		//db:      connection,
	}


	//app := iris.New()
	//app.Logger().SetLevel("debug")
	//return app
}

func Run(app *iris.Application) error {
	return app.Run(
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