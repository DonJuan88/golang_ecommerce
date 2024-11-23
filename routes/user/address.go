package routes

import (
	"github.com/gin-gonic/gin"
	usercontrollers "github.com/pintarkode/my_api/controllers/user"
)


func AddressRoutes(rg *gin.RouterGroup) {
	rg.GET("/address/:id", usercontrollers.AddressShow)
	rg.GET("/address", usercontrollers.AddressIndex)
	rg.POST("/address", usercontrollers.AddressPost)
	rg.PUT("/address/:id", usercontrollers.AddressUpdate)
	rg.DELETE("/address/:id", usercontrollers.AddressDelete)

}
