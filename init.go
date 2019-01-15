package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"web_coin/model"
)

func main() {
	db, err := gorm.Open("sqlite3", "web_coin.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.User{})
	// 为模型`User`创建表
	//db.CreateTable(&User{})

}
