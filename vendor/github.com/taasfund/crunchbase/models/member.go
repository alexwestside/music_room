package models

import (
    "github.com/jinzhu/gorm"
)

type Member struct {
    gorm.Model
    Name        int    `json:"name"`
    Position    string `json:"position"`
    Description string `json:"description"`
}
