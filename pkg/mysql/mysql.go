package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connetion DataBases
func DatabaseInit() {
	var err error
	dsn := "root:6cC6l4FXRLJfU8GWUFRB@tcp(containers-us-west-154.railway.app
:5647)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Conneted To Databases")
}
