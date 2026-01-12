package main

import (
	"fmt"

	"example/5.mysql/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	user := model.User{
		Name: "s1ark",
		Age:  24,
	}
	create(db, user)

	fmt.Println(get(db))
}

func create(db *gorm.DB, user model.User) {
	db.Create(&user)
}

func get(db *gorm.DB) model.User {
	var user model.User
	db.First(&user)
	return user
}
