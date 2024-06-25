package controllers

import (
	"ZeZeIM/database"
	"ZeZeIM/models"
	"ZeZeIM/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateGroupInput struct {
	Name    string `json:"name" binding:"required"`
	OwnerID uint   `json:"owner_id" binding:"required"`
}

func CreateGroup(c *gin.Context) {
	var input CreateGroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group := models.Group{Name: input.Name, OwnerID: input.OwnerID}
	if err := database.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func AddGroupMember(c *gin.Context) {
	var input services.AddMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddMemberToGroup(input.GroupID, input.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added to group successfully"})
}
