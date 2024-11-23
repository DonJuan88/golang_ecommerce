package admincontrollers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/helper"

	"github.com/pintarkode/my_api/models"
)

func ItemFilter(c *gin.Context) {

	var item []models.Items
	var total int64

	brands := c.Query("brands")
	category := c.Query("category")
	name := c.Query("name")
	//price := c.Query("saleprice")

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	db := config.DB

	query := db.Model(&item).Joins("Brand").Joins("Categories")

	if name != "" {
		query = query.Where("itemname like ?", "%"+name+"%")
	}
	if brands != "" {
		query = query.Where("brands.brand ?", "%"+brands+"%")
	}

	if category != "" {
		query = query.Where("category.category like ?", "&"+category+"%")
	}

	//pagination
	// Apply pagination (Offset and Limit)
	offset := (page - 1) * limit
	query.Count(&total)
	query = query.Offset(offset).Limit(limit)

	// Execute the query and load related Category and Supplier data
	query.Preload("Category").Preload("Supplier").Find(&item)
	// Eksekusi query
	if err := query.Find(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  item,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func ItemIndex(c *gin.Context) {

	var item []models.Items
	var total int64

	// Pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	//pagination
	// Apply pagination (Offset and Limit)

	// Eksekusi query

	res := config.DB.Find(&item)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  item,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// UploadFile is a function that handles the upload of a single file

func ItemPost(c *gin.Context) {

	var item *models.Items
	err := c.ShouldBind(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Check if item exists
	exists, err := helper.CheckItemExists(config.DB, item.ItemCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Item Code already registered"})
		return
	}

	res := config.DB.Create(item)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Item cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Item Created",
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

func ItemUpdate(c *gin.Context) {
	var item models.Items
	id := c.Param("id")
	err := c.ShouldBind(&item)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateItem models.Items
	res := config.DB.Model(&UpdateItem).Where("id = ?", id).Updates(item)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Item not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Item Updated",
	})
}

func ItemDelete(c *gin.Context) {
	var item models.Items
	id := c.Param("id")
	res := config.DB.Find(&item, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}
	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted",
	})

}

// correct real function
/*func UploadHandler(c *gin.Context) {

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
		return
	}

	exists, err := helper.CheckCategoryExists(config.DB, c.PostForm("category"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Category not registered"})
		return
	}

	exist, err := helper.CheckBrandExists(config.DB, c.PostForm("brand"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !exist {
		c.JSON(http.StatusConflict, gin.H{"error": "Brand not registered"})
		return
	}
	pathmaster := filepath.Dir("C:/PROJECT/FRONTEND/FLUTTER/newpawpaw/datase/")
	filename := filepath.Base(file.Filename)
	filePath := filepath.Join(pathmaster, "images", filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	fmt.Println(filePath)

	code := c.PostForm("code")
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	category := c.PostForm("category")
	brand := c.PostForm("brand")
	baseprice, err := strconv.ParseInt(c.PostForm("baseprice"), 10, 64)
	if err != nil {
		return
	}
	saleprice, err := strconv.ParseInt(c.PostForm("saleprice"), 10, 64)
	if err != nil {
		return
	}

	stock, err := strconv.ParseInt(c.PostForm("stock"), 10, 64)
	if err != nil {
		return
	}
	unit := c.PostForm("unit")
	// Simpan informasi ke database
	item := models.Items{
		ItemCode:    code,
		ItemName:    name,
		Description: desc,
		Category:    category,
		Brand:       brand,
		BasePrice:   baseprice,
		SalePrice:   saleprice,
		Stock:       stock,
		Unit:        unit,
		Image:       filePath,
	}
	fmt.Println(item)
	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
*/
