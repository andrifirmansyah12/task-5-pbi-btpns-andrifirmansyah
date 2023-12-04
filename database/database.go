package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	dsn := fmt.Sprintf(
		"%s:@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DATABASE_HOST"],
		config["DB_DATABASE"],
	)
	GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
