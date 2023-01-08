package database

import (
	"dewetour/models"
	"dewetour/pkg/mysql"
	"fmt"
	// "gorm.io/driver/mysql"
)

// Auto Migrate iff running app
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Country{}, &models.Trip{}, &models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
