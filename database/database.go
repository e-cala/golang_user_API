package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example.com/crud-user/logs"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"primaryKey"`
}

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	godotenv.Load()
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logs.Error.Println("Failed to connect database")
	}
	migrate()
}

func migrate() {
	DBConn.AutoMigrate(&User{})
	logs.Info.Println("Database migration completed!")
}
