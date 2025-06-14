package config

import (
	"final-project/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:asmonta102@tcp(127.0.0.1:3306)/asset_management_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database")
	}

	database.AutoMigrate(
		&entity.User{},
		&entity.Item{},
		&entity.Activity{},
	)

	DB = database
	fmt.Println("Database terkoneksi")
}
