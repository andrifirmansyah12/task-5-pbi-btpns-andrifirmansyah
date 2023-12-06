package controllers

import (
	"log"

	"github.com/andrifirmansyah12/projectGo/database"
	"github.com/andrifirmansyah12/projectGo/models"
	"github.com/gin-gonic/gin"
)

func FindPhotos(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Mengambil semua foto berdasarkan user_id
	var photos []models.Photo
	if err := database.GlobalDB.Where("user_id = ?", userID).Find(&photos).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Could not retrieve user photos",
		})
		return
	}

	// Menghapus password dari pengguna dan foto terkait
	for i := range photos {
		photos[i].User.Password = ""
	}

	c.JSON(200, photos)
}

func CreatePhoto(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	var photo models.Photo
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}

	// Mengisi user_id secara otomatis
	photo.UserID = uint(userID.(int))

	err = photo.CreatePhotoRecord()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating Photo",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Created Photo",
		"Data":    photo,
	})
}

func FindPhoto(c *gin.Context) {
	// Mendapatkan ID foto dari URL atau parameter request
	photoID := c.Param("id")

	// Mencari foto berdasarkan ID di database
	var photo models.Photo
	if err := database.GlobalDB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "Photo not found",
		})
		return
	}

	// Menghapus password dari pengguna terkait (opsional)
	photo.User.Password = ""

	c.JSON(200, photo)
}

func UpdatePhoto(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Mendapatkan ID foto dari URL atau parameter request
	photoID := c.Param("id")

	// Mencari foto berdasarkan ID di database
	var photo models.Photo
	if err := database.GlobalDB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "Photo not found",
		})
		return
	}

	// Bind data pembaruan dari JSON payload
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid request payload",
		})
		return
	}

	// Mengisi user_id secara otomatis
	photo.UserID = uint(userID.(int))

	// Melakukan pembaruan di database
	if err := database.GlobalDB.Save(&photo).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Error updating photo",
		})
		return
	}

	// Menghapus password dari pengguna terkait (opsional)
	photo.User.Password = ""

	c.JSON(200, gin.H{
		"Message": "Photo updated successfully",
		"Data":    photo,
	})
}

func DeletePhoto(c *gin.Context) {
	// Mendapatkan ID foto dari URL atau parameter request
	photoID := c.Param("id")

	// Mencari foto berdasarkan ID di database
	var photo models.Photo
	if err := database.GlobalDB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "Photo not found",
		})
		return
	}

	// Melakukan penghapusan di database
	if err := database.GlobalDB.Delete(&photo).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Error deleting photo",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Photo deleted successfully",
	})
}
