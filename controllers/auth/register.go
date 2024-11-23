package authlogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"
	"github.com/pintarkode/my_api/models"
)



func UserRegistration(c *gin.Context) {
	var register models.Register

	if err := c.ShouldBindJSON(&register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Complete Your Field",
		})
		return
	}

	if register.Password != register.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Match",
		})
		return

	}

	// Check if email exists
	exists, err := helper.CheckEmailExists(config.DB, register.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	paswordHash, err := helper.HashPassword(register.Password )
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

	uuid := uuid.New()
	uuidString := uuid.String()

	account := models.Accounts{
		Uuid:         uuidString,
		FirstName:    register.FirstName,
		LastName:     register.LastName,
		AccountEmail: register.Email,
		Password:     paswordHash,
		IsAdmin:      register.IsAdmin,
	}

	//fmt.Println(account)

	if err := config.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create Account",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Account Created Successfully",
	})
}
