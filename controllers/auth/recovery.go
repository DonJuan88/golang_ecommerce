package authlogin

/*import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"
	"github.com/pintarkode/my_api/models"
)

var jwtKey = []byte(config.ENV.TOKEN_LOGIN)


type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func RequestPasswordReset(c *gin.Context) {
	var input models.EmailPasswordReset
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Find user by email
	user, err := FindUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create JWT token for password reset
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create reset token"})
		return
	}

	// Send email with the reset token
	resetURL := "https://localhost:8080/api/reset-password?token=" + tokenString
	helper.SendPasswordResetEmail(input.Email, resetURL)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}


func FindUserByEmail(email string) (*models.Accounts, error) {
    var account models.Accounts
    if err := config.DB.Where("email = ?", email).First(&account).Error; err != nil {
        return nil, errors.New("user not found")
    }
    return &account, nil
}



func ResetPassword(c *gin.Context) {

	var resetPassword models.ResetPassword
    if err := c.ShouldBindJSON(&resetPassword); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Parse JWT token
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(resetPassword.Token, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    // Hash the new password
	paswordHash, err := helper.HashPassword(resetPassword.NewPassword )
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

    // Update user password in database
	var account models.Accounts
	config.DB.Model(&account).Where("id = ?",claims.Id ).Update("Password", paswordHash)

	c.JSON(http.StatusOK, gin.H{
		"message": "Password Updated",
	})
    c.JSON(http.StatusOK, gin.H{"message": "Password has been reset successfully"})
}

*/