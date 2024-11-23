package usercontrollers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pintarkode/my_api/config"

	"github.com/pintarkode/my_api/models"
)

func ItemIndex(c *gin.Context) {
	var item []models.Items
	brands := c.Query("brand")
	category := c.Query("category")
	saleprice := c.Query("saleprice")

	query := config.DB

	if category != "" {
		query = query.Where("category = ?", category)
	}

	fmt.Println("Part 1")

	// Eksekusi query
	if err := query.Find(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if brands != "" {
		query = query.Where("brand = ?", brands)
	}

	if saleprice != "" {
		query = query.Order("saleprice DESC")
	}

	sortBy := c.DefaultQuery("sort_by", "created_at") // Default sorting by created_at
	order := c.DefaultQuery("order", "asc")           // Default order is ascending

	// Validate sortBy to prevent SQL injection

	fmt.Println("Part 2")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	fmt.Println("Part 3")

	// Prevent page and pageSize from being negative or zero
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	fmt.Println("Part 4")

	// Calculate offset (skip records based on current page)
	offset := (page - 1) * pageSize

	fmt.Println("Part 5")

	// Query database with pagination
	config.DB.Offset(offset).Limit(pageSize).Find(&item)

	fmt.Println("Part 6")

	// Get total count of products (for frontend to calculate total pages)
	var total int64
	config.DB.Model(&models.Items{}).Count(&total)

	fmt.Println("Part 7")

	// Apply sorting with GORM
	config.DB.Order(fmt.Sprintf("%s %s", sortBy, order)).Offset(offset).Limit(pageSize).Find(&item)

	fmt.Println("Part 8")

	// Eksekusi query
	if err := query.Find(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        item,
		"page":        page,
		"page_size":   pageSize,
		"total":       total,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize), // Total pages
	})
}

func ItemShow(c *gin.Context) {
	var item models.Items
	id := c.Param("id")
	res := config.DB.Find(&item, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func ItemShowWithCategory(c *gin.Context) {
	var item models.Items
	cat := c.Param("category")

	res := config.DB.Find(&item, cat)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func ItemShowWithBrand(c *gin.Context) {
	var item models.Items
	brand := c.Param("brand")

	res := config.DB.Where("active =? and brand=?", true, brand).Find(&item)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}
