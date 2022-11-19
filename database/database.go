package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example.com/crud-user/logs"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error.Panic("Failed to connect database")
	}
	logs.Info.Println("Database connection successfully opened")
	DBConn.AutoMigrate(&User{})
	logs.Info.Println("Database Migrated")
}
