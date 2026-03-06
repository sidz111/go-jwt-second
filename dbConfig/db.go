package dbconfig

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(db *gorm.DB) error {
	user_name := "root"
	user_password := "root"
	host := "localhost"
	port := "3306"
	db_name := "jwt_db_user"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user_name, user_password, host, port, db_name)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = database
	return nil
}
