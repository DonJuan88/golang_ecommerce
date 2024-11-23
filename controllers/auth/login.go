package authlogin

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pintarkode/my_api/config"

	"github.com/pintarkode/my_api/models"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *gin.Context) {
	var loginDetails models.Login

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login details"})
		return
	}

	//  fmt.Println("step 1")

	var account models.Accounts
	if err := config.DB.Where("account_email = ? ", loginDetails.Email).First(&account).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// fmt.Println("step 2")
	// Compare password
	comparePasswod := account.Password //+ string(config.ENV.TOKEN_LOGIN)
	if err := bcrypt.CompareHashAndPassword([]byte(comparePasswod), []byte(loginDetails.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	//	fmt.Println("step 3")

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   account.Uuid,
		"email": account.AccountEmail,
		"exp":   time.Now().Add(time.Minute * 200).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("TOKEN_LOGIN"))) // ---ASLI---

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	//	fmt.Println("step 4")

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Author", token, 3*3600, "", "", false, true)

	c.JSON(200, gin.H{
		"Status": "login success",
		"token":  token,
	})

}

func Validate(c *gin.Context) {
	account, _ := c.Get("account")

	c.JSON(200, gin.H{
		"data": account,
	})

}

func CheckCookie(c *gin.Context) {
	// Mendapatkan semua cookie dari request
	cookies := c.Request.Cookies()
	for _, cookie := range cookies {
		// Memeriksa nama dan domain cookie
		c.String(http.StatusOK, "Cookie Name: %s, Domain: %s\n", cookie.Name, cookie.Domain)
	}
}
