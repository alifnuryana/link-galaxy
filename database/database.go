package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alifnuryana/link-galaxy/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("error load .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to database")
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("error migrate database")
	}
}
