package usercontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func AddressIndex(c *gin.Context) {
	var address []models.Addresses
	res := config.DB.Find(&address)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": address,
	})
}

func AddressPost(c *gin.Context) {
	var address *models.Addresses
	err := c.ShouldBind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := config.DB.Create(address)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Address cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New Address Created",
	})
}

func AddressShow(c *gin.Context) {
	var address models.Addresses
	uid := c.Param("uid")
	res := config.DB.Find(&address, uid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Address not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": address,
	})
}

func AddressUpdate(c *gin.Context) {
	var address models.Addresses
	uid := c.Param("uid")
	err := c.ShouldBind(&address)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateAddress models.Addresses
	res := config.DB.Model(&UpdateAddress).Where("uid = ?", uid).Updates(address)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Address not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Address Updated",
	})
}

func AddressDelete(c *gin.Context) {
	var address models.Addresses
	uid := c.Param("uid")
	res := config.DB.Find(&address, uid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Address not found",
		})
		return
	}
	config.DB.Delete(&address)
	c.JSON(http.StatusOK, gin.H{
		"message": "Address deleted",
	})
}
