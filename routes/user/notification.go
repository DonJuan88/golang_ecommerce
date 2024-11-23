package routes

import (
	"github.com/gin-gonic/gin"
	usercontrollers "github.com/pintarkode/my_api/controllers/user"
)

func UserNotifications(rg *gin.RouterGroup) {
	rg.GET("/notifications/", usercontrollers.NotificationIndex)
	rg.GET("/notifications/:id", usercontrollers.NotificationShow)
	rg.POST("/notifications", usercontrollers.NotificationPost)
}