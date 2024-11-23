package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pintarkode/my_api/config"

	"github.com/pintarkode/my_api/models"
)

func CheckAuth(c *gin.Context) {

	authHeader, _ := c.Cookie("Author")
	//ojo diganti
	//	fmt.Println("part 1 jwt")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//	fmt.Println(authHeader)
	//	fmt.Println("part 2 jwt")

	// nduwur ojo diganti

	token, _ := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_LOGIN")), nil
	})

	//	fmt.Println("part 3 jwt")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var account models.Accounts
		config.DB.First(&account, claims["sub"])

		if account.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		c.Set("account", account)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

/*
func CheckAuthAdmin(c *gin.Context) {

	authHeader, _ := c.Cookie("Author")
	//ojo diganti
	//	fmt.Println("part 1 jwt")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//	fmt.Println(authHeader)
	//	fmt.Println("part 2 jwt")

	// nduwur ojo diganti

	token, _ := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_LOGIN")), nil
	})

	//	fmt.Println("part 3 jwt")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var account models.Accounts
		config.DB.First(&account, claims["sub"])

		if account.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		c.Set("account", account)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
*/
