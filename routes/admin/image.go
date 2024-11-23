package routes

import (
	"github.com/gin-gonic/gin"
	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func ImageRoutes(rg *gin.RouterGroup){
		rg.GET("/images/:code", admincontrollers.ImageShow)
		rg.POST("/images", admincontrollers.ImagePost)
		rg.DELETE("/images/:id", admincontrollers.ImageDelete)

}