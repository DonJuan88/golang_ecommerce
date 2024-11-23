package routes

import (
	"github.com/gin-gonic/gin"

	authlogin "github.com/pintarkode/my_api/controllers/auth"
)

func AccountRoutes(rg *gin.RouterGroup){
	
		rg.GET("/account/:id", authlogin.MyAccount)
		rg.PATCH("/account/:id", authlogin.UpdateAccount)
}