package routes

import (
	"github.com/gin-gonic/gin"
	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func AdminNotifications(rg *gin.RouterGroup){
	rg.GET("/notifications/", admincontrollers.NotificationIndex)
	rg.GET("/notifications/:id", admincontrollers.NotificationShow)
	rg.POST("/notifications", admincontrollers.NotificationPost)
}