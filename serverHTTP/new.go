package serverHTTP

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/music_room/db"
)

type ServerHTTP struct {
	DB      *gorm.DB
	App     *iris.Application
	Admin 	*admin.Admin

	Address string
}

//
//type HTTPHandler struct {
//	App     *iris.Application
//	Address string
//	//db      *gorm.DB
//}


func New() *ServerHTTP {
	return &ServerHTTP{
		DB:    db.New(),
		App:   iris.New(),
		Admin: admin.New(&qor.Config{DB: DB}),
	}


	//app := iris.New()
	//app.Logger().SetLevel("debug")
	//return app
}
