package serverHTTP

import (
	"github.com/kataras/iris"
	"github.com/qor/admin"
	"github.com/music_room/db/psql"
	"github.com/music_room/db/rds"
	"github.com/qor/qor"
)

type TypeServerHTTP struct {
	PSQL  *psql.TypePSQL
	Redis *rds.TypeRedis
	App   *iris.Application
	Admin *admin.Admin
}

var Server *TypeServerHTTP

func New() {
	Server = &TypeServerHTTP{}
	Server.PSQL = psql.New()
	Server.Redis = rds.New()
	Server.Admin = admin.New(&qor.Config{DB: Server.PSQL.NewConn()})
	Server.App = iris.New()
}
