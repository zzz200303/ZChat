package main

import (
	"ZeZeIM/config"
	"ZeZeIM/controllers"
	"ZeZeIM/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	database.InitDB()

	r := gin.Default()

	r.Static("/images", config.AppConfig.ImageStoragePath)

	api := r.Group("/api")
	{
		api.POST("/upload/image", controllers.UploadImage)
		api.POST("/group", controllers.CreateGroup)
		api.POST("/group/member", controllers.AddGroupMember)
	}

	r.Run(":8080")
}
