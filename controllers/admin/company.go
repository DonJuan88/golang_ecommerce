package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func CompanyIndex(c *gin.Context) {
	var company []models.Company
	res := config.DB.Find(&company)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": company,
	})
}

func CompanyPost(c *gin.Context) {
	var company *models.Company
	err := c.ShouldBind(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := config.DB.Create(company)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New company cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New company Created",
	})
}

func CompanyShow(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	res := config.DB.Find(&company, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "company not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": company,
	})
}

func CompanyUpdate(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	err := c.ShouldBind(&company)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var Updatecompany models.Company
	res := config.DB.Model(&Updatecompany).Where("id = ?", id).Updates(company)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "company not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "company Updated",
	})
}

func CompanyDelete(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	res := config.DB.Find(&company, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "company not found",
		})
		return
	}
	config.DB.Delete(&company)
	c.JSON(http.StatusOK, gin.H{
		"message": "company deleted",
	})
}
