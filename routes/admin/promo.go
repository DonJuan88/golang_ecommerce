package routes

import (
	"github.com/gin-gonic/gin"

	admincontrollers "github.com/pintarkode/my_api/controllers/admin"
)

func PromoRoutes(rg *gin.RouterGroup){
		rg.GET("/promos/:id", admincontrollers.PromoShow)
		rg.GET("/promos", admincontrollers.PromoIndex)
		rg.POST("/promos", admincontrollers.PromoPost)
		rg.PUT("/promos/:id", admincontrollers.PromoUpdate)
		rg.DELETE("/promos/:id", admincontrollers.PromoDelete)
}