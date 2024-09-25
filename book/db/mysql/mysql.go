package mysql

import (
	"book/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL() error {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/zk2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(new(model.Book), new(model.Order), new(model.User), new(model.OrderBook))
}
