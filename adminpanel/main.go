package adminpanel

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/qor/admin"
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

func New() *http.ServeMux {
	//DB, _ := gorm.Open("sqlite3", "demo.db")
	//DB.AutoMigrate(&User{}, &Product{})
	//
	//Admin := admin.New(&admin.AdminConfig{DB: DB})
	//Admin.AddResource(&User{})
	//Admin.AddResource(&Product{})

	Admin := admin.New()
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	return mux
}