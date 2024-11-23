package admincontrollers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pintarkode/my_api/config"
	"github.com/pintarkode/my_api/models"
)

func PromoIndex(c *gin.Context) {
	var Promos []models.Promotions

	res := config.DB.Find(&Promos)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Promos,
	})
}

func PromoPost(c *gin.Context) {
	var Promos *models.Promotions
	err := c.ShouldBind(&Promos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

/* 	// Check if Promo exists
	exists, err := helper.CheckPromoExists(config.DB, Promos.PromoCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Promo Code already registered"})
		return
	}
 */
	res := config.DB.Create(Promos)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Promo cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo Created",
	})
}

func PromoShow(c *gin.Context) {
	var Promos models.Promotions
	id := c.Param("id")
	res := config.DB.Find(&Promos, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Promo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Promos,
	})
}

func PromoUpdate(c *gin.Context) {
	var Promos models.Promotions
	id := c.Param("id")
	err := c.ShouldBind(&Promos)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdatePromo models.Promotions
	res := config.DB.Model(&UpdatePromo).Where("id = ?", id).Updates(Promos)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Promo not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo Updated",
	})
}

func PromoDelete(c *gin.Context) {
	var Promos models.Promotions
	id := c.Param("id")
	res := config.DB.Find(&Promos, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Promo not found",
		})
		return
	}
	config.DB.Delete(&Promos)
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo deleted",
	})
}



//Detail Promo


func PromoDetIndex(c *gin.Context) {
	var Promos []models.PromotionDetails

	res := config.DB.Find(&Promos)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Promos,
	})
}

func PromoDetPost(c *gin.Context) {
	var Promos *models.PromotionDetails
	err := c.ShouldBind(&Promos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

/* 	// Check if Promo exists
	exists, err := helper.CheckPromoExists(config.DB, Promos.PromoCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Promo Code already registered"})
		return
	}
 */
	res := config.DB.Create(Promos)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Promo cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo Created",
	})
}

func PromoDetShow(c *gin.Context) {
	var Promos models.PromotionDetails
	id := c.Param("id")
	res := config.DB.Find(&Promos, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Promo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Promos,
	})
}

func PromoDetUpdate(c *gin.Context) {
	var Promos models.PromotionDetails
	id := c.Param("id")
	err := c.ShouldBind(&Promos)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdatePromo models.PromotionDetails
	res := config.DB.Model(&UpdatePromo).Where("id = ?", id).Updates(Promos)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Promo not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo Updated",
	})
}

func PromoDetailDelete(c *gin.Context) {
	var Promos models.PromotionDetails
	id := c.Param("id")
	res := config.DB.Find(&Promos, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Promo not found",
		})
		return
	}
	config.DB.Delete(&Promos)
	c.JSON(http.StatusOK, gin.H{
		"message": "Promo deleted",
	})
}