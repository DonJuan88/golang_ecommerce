package usercontrollers

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

func OrderPost(c *gin.Context) {
	var order *models.OrderInput
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	/* currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	ID := currentTimestamp + int64(uniqueID)

	*/
	res := config.DB.Create(order)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Order cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New Order Created",
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
	transactnoid := c.Param("transactnoid")
	err := c.ShouldBind(&order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateOrder models.Orders
	res := config.DB.Model(&UpdateOrder).Where("transactnoid = ?", transactnoid).Updates(order)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Order updated",
	})
}

func OrderDelete(c *gin.Context) {
	var order models.Orders
	transactnoid := c.Param("transactnoid")
	res := config.DB.Find(&order, transactnoid)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}
	config.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted",
	})
}
