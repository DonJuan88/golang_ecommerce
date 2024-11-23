package usercontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func PaymentIndex(c *gin.Context) {
	var payment []models.Payments

	payments := c.Param("payment_no")

	res := config.DB.Find(&payments)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("data not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

func PaymentPost(c *gin.Context) {
	var payment *models.Payments
	err := c.ShouldBind(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := config.DB.Create(payment)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Payment cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

func PaymentShow(c *gin.Context) {
	var payment models.Payments
	id := c.Param("account_id")
	res := config.DB.Find(&payment, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

func PaymentUpdate(c *gin.Context) {
	var payment models.Payments
	id := c.Param("account_id")
	err := c.ShouldBind(&payment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdatePayment models.Payments
	res := config.DB.Model(&UpdatePayment).Where("account_id = ?", id).Updates(payment)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payment not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

/*
func PaymentDelete(c *gin.Context) {
	var payment models.Payments
	id := c.Param("id")
	res := config.DB.Find(&payment, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}
	config.DB.Delete(&payment)
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment deleted",
	})
} */
