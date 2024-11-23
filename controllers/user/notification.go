package usercontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func NotificationIndex(c *gin.Context) {
	var Notifications []models.Notification

	res := config.DB.Find(&Notifications)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Notifications,
	})
}

func NotificationPost(c *gin.Context) {
	var Notifications *models.Notification
	err := c.ShouldBind(&Notifications)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	

	

	res := config.DB.Create(Notifications)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Notification cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification Created",
	})
}

func NotificationShow(c *gin.Context) {
	var Notifications models.Notification
	id := c.Param("id")
	res := config.DB.Find(&Notifications, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Notification not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Notifications,
	})
}

func NotificationUpdate(c *gin.Context) {
	var Notifications models.Notification
	id := c.Param("id")
	err := c.ShouldBind(&Notifications)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateNotification models.Notification
	res := config.DB.Model(&UpdateNotification).Where("id = ?", id).Updates(Notifications)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Notification not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification Updated",
	})
}

func NotificationDelete(c *gin.Context) {
	var Notifications models.Notification
	id := c.Param("id")
	res := config.DB.Find(&Notifications, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Notification not found",
		})
		return
	}
	config.DB.Delete(&Notifications)
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification deleted",
	})
}
