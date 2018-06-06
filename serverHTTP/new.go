package serverHTTP

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/music_room/db/psql"
	"github.com/music_room/db/rds"
	"github.com/go-redis/redis"
)

type TypeServerHTTP struct {
	PSQL  *gorm.DB
	Redis *redis.Client
	App   *iris.Application
	Admin *admin.Admin
}

var Server *TypeServerHTTP

func New(){
	Server = &TypeServerHTTP{}

	Server.PSQL = psql.New()
	Server.Redis = rds.New()
	Server.Admin = admin.New(&qor.Config{DB: Server.PSQL})
	Server.App = iris.New()
	Server.App.Logger().SetLevel("debug")

}

func (s TypeServerHTTP) F() {

}