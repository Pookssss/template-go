package database

import (
	"fmt"
	"os"
	"template-go/models"

	// "template-go/models"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORTS := os.Getenv("DB_PORTS")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORTS, DB_NAME)
	// dsn := "root:@tcp(127.0.0.1:23306)/learning"
	// dsn := "userforlearning:t@sswo9#232@tcp(127.0.0.1:23306)/learning"

	database, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println('\n')
	fmt.Println("---------------- Connected to database ----------------")
	DB = database
	DB.AutoMigrate(&models.User{})

}
