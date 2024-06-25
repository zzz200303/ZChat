package services

import (
	"ZeZeIM/database"
	"ZeZeIM/models"
)

type AddMemberInput struct {
	GroupID uint `json:"group_id" binding:"required"`
	UserID  uint `json:"user_id" binding:"required"`
}

func AddMemberToGroup(groupID uint, userID uint) error {
	groupMember := models.GroupMember{GroupID: groupID, UserID: userID}
	if err := database.DB.Create(&groupMember).Error; err != nil {
		return err
	}
	return nil
}
