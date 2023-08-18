package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

type User struct {
	ID       int
	Email    string
	Password string
}

func main() {
	// Create DB pool
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	//TxSelect()
	//TxDelete()
	//TxInsert()
	//TxUpdate()
}

func TxSelect() {
	var myUser User

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	err = tx.QueryRow("SELECT id, email, password FROM users WHERE id = $1", 1).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
	}
	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", myUser.Email, myUser.Password)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func TxDelete() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("DELETE FROM users  where id=$1")
	if err != nil {
		log.Fatal(err) // Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if result, err := stmt.Exec(user.ID); err != nil {
			log.Fatal(err)
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func TxInsert() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO users (id,email,password) VALUES($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if result, err := stmt.Exec(user.ID, user.Email, user.Password); err != nil {
			log.Fatal(err)
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func TxUpdate() {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal("Failed to begin tx: ", err)
	}

	//预要通过主键更改到数据库里
	var user User = User{ID: 3, Email: "dong@qq.com", Password: "abcdedf120"}
	//执行更改操作
	_, err = tx.Exec("UPDATE  users SET email=$1, password=$2 where id=$3", user.Email, user.Password, user.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	//打印日志
	log.Printf("update ok!!!")
}
