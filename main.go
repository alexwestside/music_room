package main

import (
	"github.com/music_room/adminpanel"
	"github.com/music_room/api"
	"flag"
	"os"
	"github.com/music_room/db"
	"github.com/music_room/music_room_server"
	"github.com/jinzhu/gorm"
	"fmt"
)

type argv struct {
	host    string
	port    string
	migrate bool
}

func migrate(connection *gorm.DB) {
	fmt.Println("Migrate")

	//TODO: ADD CONSTRAINTS
	//connection.AutoMigrate(&models.Project{}, &models.Investment{}, &models.Member{})
	//
	//fmt.Println("Migrations done")
}



func parseArgs() *argv {
	var flags argv

	helpHost := `Specify host`
	flag.StringVar(&flags.host, "host", "0.0.0.0", helpHost)

	helpPort := `Specify port to listen requests`
	flag.StringVar(&flags.port, "port", "80", helpPort)

	helpMigrate := `Migrate`
	flag.BoolVar(&flags.migrate, "migrate", false, helpMigrate)

	flag.Parse()

	return &flags
}

func main() {

	go adminpanel.Run()

	flags := parseArgs()

	connection, errConnDB := db.New(
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DB"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
	)
	if errConnDB != nil {
		panic(errConnDB)
	}
	defer connection.Close()

	if flags.migrate == true {
		migrate(connection)
	}

	server := music_room_server.New(flags.host, flags.port, connection)

	api.Routing(server.App)

	errServer := music_room_server.Run(server.App)
	if errServer != nil {
		panic(errServer)
	}
}
