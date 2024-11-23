package routes

import (
	"github.com/gin-gonic/gin"

	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func ItemRoutes(rg *gin.RouterGroup){
		rg.GET("/items/:id", admincontrollers.ItemShow)
		rg.GET("/items", admincontrollers.ItemIndex)
		rg.POST("/items", admincontrollers.ItemPost)
		rg.PUT("/items/:id", admincontrollers.ItemUpdate)
		rg.DELETE("/items/:id", admincontrollers.ItemDelete)


}