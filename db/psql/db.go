package psql

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func makeMigrations(connection *gorm.DB) {
	migrate := os.Getenv("MIGRATE")

	if migrate == "1" {
		fmt.Println("Migrate")

		//TODO: ADD CONSTRAINTS
		connection.AutoMigrate()

		fmt.Println("Migrations done")
	}
}


func New() *gorm.DB {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	pgname := os.Getenv("PGNAME")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASS")

	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		host, port, pgname, user, pass)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	makeMigrations(connection)

	return connection
}
