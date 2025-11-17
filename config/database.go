package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"stock/models"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/stock_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}

	db.AutoMigrate(&models.Item{})

	fmt.Println("✅ Database connected")

	DB = db
}
