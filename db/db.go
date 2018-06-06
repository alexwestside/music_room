package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"flag"
	"github.com/taasfund/crunchbase/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type argv struct {
	host    string
	port    string
	migrate bool
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


func makeMigrations(connection *gorm.DB) {
	migrate := os.Getenv("MIGRATE")

	if migrate == "1" {
		fmt.Println("Migrate")

		//TODO: ADD CONSTRAINTS
		connection.AutoMigrate(&models.Project{}, &models.Investment{}, &models.Member{})

		fmt.Println("Migrations done")
	}
}


func New() *gorm.DB {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	db := os.Getenv("DB")
	login := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")

	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		host, port, db, login, pass)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	makeMigrations(connection)

	defer connection.Close()

	return connection
}
