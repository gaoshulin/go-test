package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/galen")
	if err != nil {
		fmt.Println("open mysql failed.", err)
		return
	}

	Db = database
}

func main() {
	var query string

	// insert
	query = "insert into person(username, sex, email)values(? ,?, ?)"
	res, err := Db.Exec(query, "galen", "male", "galen@qq.com")
	if err != nil {
		fmt.Println("insert failed.", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("exec failed.", err)
		return
	}
	fmt.Println("insert id:", id)

	// update
	query = "update person set username = ? where user_id = ?"
	res, err = Db.Exec(query, "later", 3)
	if err != nil {
		fmt.Println("update failed.", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update success:", row)

	// delete
	query = "delete from person where user_id = ?"
	res, err = Db.Exec(query, 2)
	if err != nil {
		fmt.Println("delete failed.", err)
		return
	}
	row, err = res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed. ", err)
	}
	fmt.Println("delete success: ", row)

	// select
	var person []Person
	query = "select * from person order by user_id desc"
	err = Db.Select(&person, query)
	if err != nil {
		fmt.Println("select failed.", err)
		return
	}
	fmt.Println("select success:", person)

	// mysql 事务
	//coon, err := Db.Begin()
	//coon.Rollback()
	//coon.Commit()

}
