package main

import (
	"fmt"

	"example/5.mysql/model"
	"example/5.mysql/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//init
	//createTable(db)

	insert(db, "tom", 18)

	fmt.Println(get(db, 1))
}

func createTable(db *gorm.DB) {
	db.Exec(sql.Create)
}

func insert(db *gorm.DB, name string, age int) {
	db.Exec(sql.Insert, name, age)
}

func get(db *gorm.DB, id int64) model.User {
	var result model.User
	db.Raw(sql.Select, id).Scan(&result)
	return result
}

func delete(db *gorm.DB, id int64) {
	db.Exec(sql.Delete, id)
}

func update(db *gorm.DB, name string, age int, id int64) {
	db.Exec(sql.Update, name, age, id)
}
