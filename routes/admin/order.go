package routes

import (
	"github.com/gin-gonic/gin"
	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func OrderRoutes(rg *gin.RouterGroup){
	
		rg.GET("/orders", admincontrollers.OrderIndex)
		rg.PUT("/orders/:id", admincontrollers.OrderUpdate)
		rg.GET("/orders/:id", admincontrollers.OrderShow)

}