package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func OrderIndex(c *gin.Context) {
	var order []models.Orders
	res := config.DB.Find(&order)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("data not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func OrderShow(c *gin.Context) {
	var order models.Orders
	id := c.Param("id")
	res := config.DB.Find(&order, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func OrdersShow(c *gin.Context) {
	var order models.Orders
	accid := c.Param("accid")
	res := config.DB.Find(&order, accid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func OrderUpdate(c *gin.Context) {
	var order models.Orders
	id := c.Param("id")
	err := c.ShouldBind(&order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateOrder models.Orders
	res := config.DB.Model(&UpdateOrder).Where("id = ?", id).Updates(order)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}
