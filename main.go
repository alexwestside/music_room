package main

import (
	"github.com/music_room/serverHTTP"
	"github.com/music_room/adminpanel"
	"github.com/music_room/application"
	"time"
	"fmt"
)


func main() {
	time.Sleep(time.Second * 35)
	fmt.Println("Wait for PSQL...")

	serverHTTP.New()

	go adminpanel.Run()

	application.Run()
}
