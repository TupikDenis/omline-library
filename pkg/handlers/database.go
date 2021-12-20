package handlers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3307)/yana?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
