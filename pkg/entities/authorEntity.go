package entities

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Id         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorName string `json:"authorName"`
}
