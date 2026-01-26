package main

import (
	gsql "database/sql"
	"fmt"

	"example/5.mysql/model"
	"example/5.mysql/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gsql.Open("mysql", dsn)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	//createTable_(db)
	fmt.Println(get_(db, 1))
}

func createTable_(db *gsql.DB) {
	db.Exec(sql.Create)
}

func insert_(db *gsql.DB, name string, age int) {
	db.Exec(sql.Insert, name, age)
}

func get_(db *gsql.DB, id int64) model.User {
	var result model.User
	db.QueryRow(sql.Select, id).Scan(&result.ID, &result.Name, &result.Age, &result.CreatedAt, &result.UpdatedAt)
	return result
}

func delete_(db *gsql.DB, id int64) {
	db.Exec(sql.Delete, id)
}

func update_(db *gsql.DB, name string, age int, id int64) {
	db.Exec(sql.Update, name, age, id)
}
