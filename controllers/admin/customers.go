package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func CustomerIndex(c *gin.Context) {
	var account []models.Accounts

	res := config.DB.Where("is_admin = ?", "FALSE").Find(&account)
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

/*
	func AccountPost(c *gin.Context) {
		var account *models.Accounts
		err := c.ShouldBind(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		res := config.DB.Create(account)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "New Account cannot created",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Account Created",
		})
	}
*/
func CustomerShow(c *gin.Context) {
	var account models.Accounts
	id := c.Param("uid")
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

/*
func AccountUpdate(c *gin.Context) {
	var account models.Accounts
	id := c.Param("id")
	err := c.ShouldBind(&account)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateAccount models.Accounts
	res := config.DB.Model(&UpdateAccount).Where("id = ?", id).Updates(account)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Account not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Account Updated",
	})
}

func AccountDelete(c *gin.Context) {
	var account models.Accounts
	id := c.Param("id")
	res := config.DB.Find(&account, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found",
		})
		return
	}
	config.DB.Delete(&account)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted",
	})
}*/
