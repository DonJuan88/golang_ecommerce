package routes

import (
	"github.com/gin-gonic/gin"
	authlogin "github.com/pintarkode/my_api/controllers/auth"
)


func LoginRoutes(rg *gin.RouterGroup){
	
	rg.POST("/login", authlogin.UserLogin)
	rg.POST("/register", authlogin.UserRegistration)
	rg.GET("/validate", authlogin.Validate)
	rg.POST("/logout", authlogin.Logout)
//	rg.POST("/password-reset", authlogin.RequestPasswordReset)
//	rg.POST("/password-reset-confirm", authlogin.ResetPassword)
}