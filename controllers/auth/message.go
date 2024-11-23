package authlogin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)



func MessageIndex(c *gin.Context) {
	var message []models.Message
	res := config.DB.Find(&message)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": message,
	})
}

func MessagePost(c *gin.Context) {
	var Message *models.Message
	err := c.ShouldBind(&Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := config.DB.Create(Message)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Message cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New Message Created",
	})
}

func MessageShow(c *gin.Context) {
	var Message models.Message
	uid := c.Param("uid")
	res := config.DB.Find(&Message, uid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Message not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Message,
	})
}

func MessageUpdate(c *gin.Context) {
	var Message models.Message
	uid := c.Param("uid")
	err := c.ShouldBind(&Message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateMessage models.Message
	res := config.DB.Model(&UpdateMessage).Where("uid = ?", uid).Updates(Message)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Message not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Message Updated",
	})
}

func MessageDelete(c *gin.Context) {
	var Message models.Message
	uid := c.Param("uid")
	res := config.DB.Find(&Message, uid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Message not found",
		})
		return
	}
	config.DB.Delete(&Message)
	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted",
	})
}
