package routes

import (
	"github.com/gin-gonic/gin"

	authlogin "github.com/pintarkode/my_api/controllers/auth"
)


func AccountUserRoutes( rg *gin.RouterGroup){

		rg.GET("/accounts/:id", authlogin.MyAccount)
		rg.PUT("/accounts/:id", authlogin.UpdateAccount)
		rg.PATCH("/account/:id", authlogin.UpdateAccount)

}

