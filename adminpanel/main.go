package adminpanel

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/music_room/serverHTTP"
)

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}



func Run() {
	//DB, _ := gorm.Open("sqlite3", "demo.db")
	//DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	//Admin := admin.New(&qor.Config{DB: DB})

	// Allow to use Admin to manage User, Product
	serverHTTP.Server.Admin.AddResource(&User{})
	serverHTTP.Server.Admin.AddResource(&Product{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount adminpanel interface to mux
	serverHTTP.Server.Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}