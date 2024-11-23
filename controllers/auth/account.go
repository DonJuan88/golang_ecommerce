package authlogin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"
	"github.com/pintarkode/my_api/models"
	"golang.org/x/crypto/bcrypt"
)

/* func AccountIndex(c *gin.Context) {
	var account []models.Accounts
	res := config.DB.Find(&account)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("account not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})
}
*/

func MyAccount(c *gin.Context) {
	var account models.Accounts
	id := c.Param("id")
	res := config.DB.Find(&account, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})
}

func UpdateAccount(c *gin.Context) {
	var account models.Accounts
	id := c.Param("id")
	// Cari user berdasarkan ID
	if err := config.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var UpdateAccount models.UpdateAccount

	if err := c.ShouldBindJSON(&UpdateAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(UpdateAccount.LastPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	if UpdateAccount.Password != UpdateAccount.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Match",
		})
		return

	}

	paswordHash, err := helper.HashPassword(UpdateAccount.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

	//res := config.DB.Model(&UpdateAccount).Where("id = ?", id).Updates(account)
	config.DB.Model(&account).Where("id = ?", id).Update("Password", paswordHash)

	c.JSON(http.StatusOK, gin.H{
		"message": "Password Updated",
	})
}
