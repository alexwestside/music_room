package models

import (
    "github.com/jinzhu/gorm"
)

type Investment struct {
    gorm.Model
    Amount    float64 `json:"amount"`
    Currency  string  `gorm:"size:10" json:"currency"` //better to separate this entity into another table in the future
    Date      int     `gorm:"type:timestamp" json:"date"`
    ProjectID uint64  `gorm:"type:bigint REFERENCES projects(id)""json:"project_id"`
}
