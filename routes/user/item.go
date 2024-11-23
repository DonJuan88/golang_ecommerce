package routes

import (
	"github.com/gin-gonic/gin"
	usercontrollers "github.com/pintarkode/my_api/controllers/user"
)

func ItemIndexRoutes(rg *gin.RouterGroup){
	rg.GET("/items/:id", usercontrollers.ItemShow)
	rg.GET("/items", usercontrollers.ItemIndex)
	rg.GET("/items/brands/:brand", usercontrollers.ItemShowWithBrand)
}