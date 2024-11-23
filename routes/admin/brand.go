package routes

import (
	"github.com/gin-gonic/gin"

	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func BrandRoutes(rg *gin.RouterGroup){
	
		rg.GET("/brands/:id", admincontrollers.BrandShow)
		rg.GET("/brands", admincontrollers.BrandIndex)
		rg.POST("/brands", admincontrollers.BrandPost)
		rg.PUT("/brands/:id", admincontrollers.BrandUpdate)
		rg.DELETE("/brands/:id", admincontrollers.BrandDelete)
}