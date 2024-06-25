package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name     string `gorm:"unique"`
	OwnerID  uint
	Members  []User    `gorm:"many2many:group_members;"`
	Messages []Message `gorm:"polymorphic:Messageable;"`
}
