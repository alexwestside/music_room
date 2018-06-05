package main

import (
	"github.com/music_room/adminpanel"
	"github.com/music_room/serverHTTP"
	"github.com/jinzhu/gorm"
	"fmt"
)







func main() {

	go adminpanel.Run()

	serverHTTP.New()


	//
	//connection, errConnDB := db.New(
	//	os.Getenv("DBHOST"),
	//	os.Getenv("DBPORT"),
	//	os.Getenv("DB"),
	//	os.Getenv("DBUSER"),
	//	os.Getenv("DBPASS"),
	//)
	//if errConnDB != nil {
	//	panic(errConnDB)
	//}
	//defer connection.Close()

	//if flags.migrate == true {
	//	migrate(connection)
	//}
	//
	//server := serverHTTP.New(flags.host, flags.port, connection)
	//
	//api.Routing(server.App)
	//
	//errServer := serverHTTP.Run(server.App)
	//if errServer != nil {
	//	panic(errServer)
	//}
}
