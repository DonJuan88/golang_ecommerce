package authlogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	authHeader, _ := c.Cookie("Author")

	//ojo diganti

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.SetCookie("Author", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged out successfully",
	})
}
