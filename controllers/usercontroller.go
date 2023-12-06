package controllers

import (
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

func UpdateUsers(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Mencari foto berdasarkan ID di database
	var user models.User
	if err := database.GlobalDB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "Users not found",
		})
		return
	}

	// Bind data pembaruan dari JSON payload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid request payload",
		})
		return
	}

	// Hashing password
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(500, gin.H{
			"Error": "Error Hashing Password",
		})
		c.Abort()
		return
	}

	// Melakukan pembaruan di database
	if err := database.GlobalDB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Error updating users",
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"Message": "Users updated successfully",
		"Data":    user,
	})
}

func DeleteUsers(c *gin.Context) {
	// Mendapatkan ID foto dari URL atau parameter request
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"Error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Mencari foto berdasarkan ID di database
	var user models.User
	if err := database.GlobalDB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "Users not found",
		})
		return
	}

	// Melakukan penghapusan di database
	if err := database.GlobalDB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"Error": "Error deleting users",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Users deleted successfully",
	})
}
