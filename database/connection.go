package database

import (
	"template-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// user:pass@tcp(127.0.0.1:3306)
	// dsn := "userforlearning:t@sswo9#232@tcp(10.0.1.11:3306)/learning"
	dsn := "root:@tcp(127.0.0.1:3306)/learning"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = database

	database.AutoMigrate(&models.User{})

}
