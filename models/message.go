package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID        uint
	ReceiverID      uint
	Content         string
	ImageURL        string
	MessageableID   uint
	MessageableType string
}
