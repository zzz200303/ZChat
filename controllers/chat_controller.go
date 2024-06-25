package controllers

import (
	"ZeZeIM/database"
	"ZeZeIM/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendMessageInput struct {
	SenderID uint   `json:"sender_id" binding:"required"`
	GroupID  uint   `json:"group_id" binding:"required"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}

func SendMessage(c *gin.Context) {
	var input SendMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := models.Message{
		SenderID:        input.SenderID,
		Content:         input.Content,
		ImageURL:        input.ImageURL,
		MessageableID:   input.GroupID,
		MessageableType: "Group",
	}

	if err := database.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 发送消息给群组中的所有成员（未实现WebSocket部分）
	// services.BroadcastMessageToGroup(input.GroupID, message)

	c.JSON(http.StatusOK, message)
}
