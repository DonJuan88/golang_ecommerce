package routes

import (
	"github.com/gin-gonic/gin"
	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func CustomerRoutes(rg *gin.RouterGroup){
	rg.GET("/customers/:id", admincontrollers.CustomerShow)
	rg.GET("/customers", admincontrollers.CustomerIndex)

}