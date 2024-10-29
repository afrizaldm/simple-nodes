package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(500);not null" json:"name"`
	ParentID uint   `gorm:"" json:"parent_id"`
}
