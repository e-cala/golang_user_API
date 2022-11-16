package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
	DBConn.AutoMigrate(&User{})
	fmt.Println("Database Migrated")
}
