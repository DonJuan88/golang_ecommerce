package routes

import (
	"github.com/gin-gonic/gin"

	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func CategoryRoutes(rg *gin.RouterGroup){
		rg.GET("/categories/:id", admincontrollers.CategoryShow)
		rg.GET("/categories", admincontrollers.CategoryIndex)
		rg.POST("/categories", admincontrollers.CategoryPost)
		rg.PUT("/categories/:id", admincontrollers.CategoryUpdate)
		rg.DELETE("/categories/:id", admincontrollers.CategoryDelete)
}