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

	//createTable(db)
	//update(db, "s1s1", 28, 1)
	//insert(db, "s2s2", 29)
	//delete(db, 2)

	fmt.Println(get(db, 1))
}

func createTable(db *gsql.DB) {
	db.Exec(sql.Create)
}

func insert(db *gsql.DB, name string, age int) {
	db.Exec(sql.Insert, name, age)
}

func get(db *gsql.DB, id int64) model.User {
	var result model.User
	db.QueryRow(sql.Select, id).Scan(&result.ID, &result.Name, &result.Age, &result.CreatedAt, &result.UpdatedAt)
	return result
}

func delete(db *gsql.DB, id int64) {
	db.Exec(sql.Delete, id)
}

func update(db *gsql.DB, name string, age int, id int64) {
	db.Exec(sql.Update, name, age, id)
}
