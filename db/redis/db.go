package redis

import (
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
)

func New() *gorm.DB {
	host := os.Getenv("RDSHOST")
	port := os.Getenv("RDSORT")
	pgname := os.Getenv("RDSNAME")

	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s",
		host, port, pgname)

	connection, err := gorm.Open("redis", connStr)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	return connection
}