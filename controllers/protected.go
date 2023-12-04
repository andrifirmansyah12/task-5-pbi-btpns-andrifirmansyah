package controllers

import (
	"log"
	"net/http"

	"github.com/andrifirmansyah12/projectGo/database"
	"github.com/andrifirmansyah12/projectGo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Profile(c *gin.Context) {
	var user models.User
	user_id, _ := c.Get("user_id")
	result := database.GlobalDB.Where("id = ?", user_id).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}
	user.Password = ""
	c.JSON(200, user)
}

// ======================================= CRUD Photo =========================================

func FindPhotos(c *gin.Context) {
	var user models.User
	user_id, _ := c.Get("user_id")

	// Ambil user berdasarkan user_id
	result := database.GlobalDB.Where("id = ?", user_id).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		return
	}
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		return
	}

	// Preload untuk mengambil relasi Photos
	var photos []models.Photo
	if err := database.GlobalDB.Preload("Photos").Find(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Photos",
		})
		return
	}

	// Menghapus password dari user dan foto terkait
	user.Password = ""
	for i := range photos {
		photos[i].User.Password = ""
	}

	// Menggabungkan data user dengan foto terkait
	responseData := gin.H{
		"user":   user,
		"photos": photos,
	}

	c.JSON(200, responseData)
}

func CreatePhoto(c *gin.Context) {
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
	})
}

func FindPhoto(c *gin.Context) {
	var photos models.Photo

	db := c.MustGet("err").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {

	db := c.MustGet("err").(*gorm.DB)
	var task models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var input models.UpdatePhotoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Photo
	updatedInput.Title = input.Title
	updatedInput.Caption = input.Caption
	updatedInput.PhotoURL = input.PhotoURL
	// updatedInput.UserID = input.UserID

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeletePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Photo
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
