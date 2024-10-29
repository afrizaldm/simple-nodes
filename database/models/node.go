package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(500);not null"`
	ParentID uint   `gorm:"index"`
}
