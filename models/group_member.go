package models

import (
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupID uint
	UserID  uint
}
