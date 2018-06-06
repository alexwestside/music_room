package serverHTTP

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/music_room/db"
)

type TypeServerHTTP struct {
	DB      *gorm.DB
	App     *iris.Application
	Admin 	*admin.Admin
}

var Server *TypeServerHTTP

func New(){
	Server = &TypeServerHTTP{}

	Server.DB = db.New()
	Server.Admin = admin.New(&qor.Config{DB: Server.DB})
	Server.App = iris.New()
	Server.App.Logger().SetLevel("debug")

}

func (s TypeServerHTTP) F() {

}