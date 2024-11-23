package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"
	"github.com/pintarkode/my_api/models"
)

func CategoryIndex(c *gin.Context) {
	var categories []models.Categories
	res := config.DB.Find(&categories)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CategoryPost(c *gin.Context) {
	var categories *models.Categories
	err := c.ShouldBind(&categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	exists, err := helper.CheckCategoryExists(config.DB, categories.CategoryCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Category Code already registered"})
		return
	}

	res := config.DB.Create(categories)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CategoryShow(c *gin.Context) {
	var categories models.Categories
	id := c.Param("id")
	res := config.DB.Find(&categories, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CategoryUpdate(c *gin.Context) {
	var categories models.Categories
	id := c.Param("id")
	err := c.ShouldBind(&categories)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateCategory models.Categories
	res := config.DB.Model(&UpdateCategory).Where("id = ?", id).Updates(categories)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func CategoryDelete(c *gin.Context) {
	var categories models.Categories
	id := c.Param("id")
	res := config.DB.Find(&categories, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}
	config.DB.Delete(&categories)
	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted",
	})
}
