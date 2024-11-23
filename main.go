package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/pintarkode/my_api/config"
	routes_u "github.com/pintarkode/my_api/routes/user"

	routes_a "github.com/pintarkode/my_api/routes/admin"
	routes_shared "github.com/pintarkode/my_api/routes/auth"
)

var DB *gorm.DB

func main() {

	config.LoadConfig()
	config.DatabaseConnection()
	//	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Starting Application...")
	config.DatabaseConnection()

	fmt.Println("Application are Ready to Use")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(cors.Default())
	
	r.SetTrustedProxies(nil)
	api := r.Group("/api")

	user := api.Group("/user")

	//user := api.Group("/user", middleware.CheckAuth)
	{
		routes_u.AccountUserRoutes(user)
		routes_u.AddressRoutes(user)
		routes_u.ItemIndexRoutes(user)
		routes_u.OrderRoutes(user)
		routes_u.PaymentRoutes(user)
		routes_shared.LoginRoutes(user)

	}


	// For admin
	admin := api.Group("/admin")

	//	admin := api.Group("/admin", middleware.CheckAuth)
	{

		routes_a.AccountRoutes(admin)
		routes_a.BrandRoutes(admin)
		routes_a.CategoryRoutes(admin)
		routes_a.CustomerRoutes(admin)
		routes_a.ImageRoutes(admin)
		routes_a.ItemRoutes(admin)
		routes_a.OrderRoutes(admin)
		routes_shared.LoginRoutes(admin)
		routes_a.CompanyRoutes(admin)
	}

	//execute
	r.ForwardedByClientIP = true
	r.Run(fmt.Sprintf(":%v", config.ENV.URL_PORT))

}
