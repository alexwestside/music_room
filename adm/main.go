package adm

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
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
	//serverHTTP.Server.PSQL.AutoMigrate(&User{}, &Product{})

	// Initalize
	//Admin := admin.New(&qor.Config{DB: DB})

	// Allow to use Admin to manage User, Product
	//serverHTTP.Server.Admin.AddResource(&User{})
	//serverHTTP.Server.Admin.AddResource(&Product{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount adm interface to mux
	//serverHTTP.Server.Admin.MountTo("/admin", mux)

	fmt.Println("AdminPanel Listening on: http://localhost:9000/admin")
	errServer := http.ListenAndServe(":9000", mux)

	if errServer != nil {
		panic(errServer)
	}
}