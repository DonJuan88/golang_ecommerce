package routes

import (
	"github.com/gin-gonic/gin"
	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func CompanyRoutes(rg *gin.RouterGroup){
	rg.GET("/company", admincontrollers.CompanyShow)
	rg.POST("/company", admincontrollers.CompanyPost)
	rg.PUT("company/:id", admincontrollers.CompanyPost)
	rg.DELETE("company/:id", admincontrollers.CompanyDelete)
}