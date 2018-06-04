package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
)


func New(host string, port string, db string, login string, pass string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		host, port, db, login, pass)

	return gorm.Open("postgres", connStr)
}