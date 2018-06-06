package models

import (
    "github.com/jinzhu/gorm"
)

type Project struct {
    gorm.Model
    Name        string       `gorm:"size:100" json:"id"`
    URL         string       `gorm:"size:1024" json:"url"`
    Description string       `gorm:"size:10240" json:"description"`
    StartDate   int64        `gorm:"type:timestamp" json:"start_date"`
    EndDate     int64        `gorm:"type:timestamp" json:"end_date"`
    Investments []Investment `gorm:"foreignkey:ProjectID;association_foreignkey:Refer" json:"investments"`
    Members     []Member     `gorm:"many2many:project_members" json:"members"`
}
