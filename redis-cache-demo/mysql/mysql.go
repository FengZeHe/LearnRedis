package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var sqlDB *gorm.DB
var db *gorm.DB

func InitMysql() (err error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect msyql error", err)
	}
	fmt.Println("mysql connect success!")
	return err
}
