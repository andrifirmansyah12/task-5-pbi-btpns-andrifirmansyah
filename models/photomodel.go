package models

import (
	"time"

	"github.com/andrifirmansyah12/projectGo/database"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"title" gorm:"unique"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	User      User
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (photo *Photo) CreatePhotoRecord() error {
	result := database.GlobalDB.Create(&photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type UpdatePhotoInput struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint
}
