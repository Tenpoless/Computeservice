package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	username := ""
	password := ""
	host := ""
	port := "3306"
	name := ""

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if errDB != nil {
		panic ("Failed to connect database")
	}

	fmt.Println("connected to database")
}

// func AutoMigrate() {
// 	DB.AutoMigrate(&model.User{}, )
// }
