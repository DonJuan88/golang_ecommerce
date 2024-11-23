package routes

import (
	"github.com/gin-gonic/gin"
	usercontrollers "github.com/pintarkode/my_api/controllers/user"
)

func OrderRoutes(rg *gin.RouterGroup){
		rg.GET("/orders", usercontrollers.OrderIndex)
		rg.POST("/orders", usercontrollers.OrderPost)
		rg.PUT("/orders/:id", usercontrollers.OrderUpdate)
		rg.DELETE("/orders/:id", usercontrollers.OrderDelete)
}