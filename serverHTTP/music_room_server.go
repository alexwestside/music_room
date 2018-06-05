package serverHTTP

import (
	"github.com/kataras/iris"
)







func Run(app *iris.Application) error {
	return app.Run(
		// Start the web serverHTTP at localhost:80
		iris.Addr(":8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err serverHTTP closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}