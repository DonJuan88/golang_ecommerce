package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"
	"github.com/pintarkode/my_api/models"
)

func BrandIndex(c *gin.Context) {
	var brands []models.Brands

	res := config.DB.Find(&brands)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": brands,
	})
}

func BrandPost(c *gin.Context) {
	var brands *models.Brands
	err := c.ShouldBind(&brands)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	// Check if brand exists
	exists, err := helper.CheckBrandExists(config.DB, brands.BrandCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Brand Code already registered"})
		return
	}

	res := config.DB.Create(brands)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Brand cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Brand Created",
	})
}

func BrandShow(c *gin.Context) {
	var brands models.Brands
	id := c.Param("id")
	res := config.DB.Find(&brands, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Brand not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": brands,
	})
}

func BrandUpdate(c *gin.Context) {
	var brands models.Brands
	id := c.Param("id")
	err := c.ShouldBind(&brands)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateBrand models.Brands
	res := config.DB.Model(&UpdateBrand).Where("id = ?", id).Updates(brands)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Brand not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Brand Updated",
	})
}

func BrandDelete(c *gin.Context) {
	var brands models.Brands
	id := c.Param("id")
	res := config.DB.Find(&brands, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Brand not found",
		})
		return
	}
	config.DB.Delete(&brands)
	c.JSON(http.StatusOK, gin.H{
		"message": "Brand deleted",
	})
}
