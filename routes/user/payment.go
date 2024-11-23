package routes

import (
	"github.com/gin-gonic/gin"
	usercontrollers "github.com/pintarkode/my_api/controllers/user"
)

func PaymentRoutes(rg *gin.RouterGroup) {
	rg.GET("/payments/:id", usercontrollers.PaymentShow)
	rg.GET("/payments", usercontrollers.PaymentIndex)
	rg.POST("/payments", usercontrollers.PaymentPost)
	rg.PUT("/payments/:id", usercontrollers.PaymentUpdate)

}