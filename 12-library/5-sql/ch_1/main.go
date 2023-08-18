package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // 注册psql数据库驱动
	"log"
)

/*
	1.演示注册数据库驱动
	2.Select/Insert/Update/Delete 操作数据
*/

const (
	DB_DSN = "host=10.2.0.105 port=5432 user=postgres password=lab_password dbname=postgres sslmode=disable"
)

var db *sql.DB
var err error

func main() {
	// 创建数据库链接
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	Select() // 查询数据
	//Insert() // 插入数据
	//Update() // 更新数据
	Delete() // 删除数据
}

type User struct {
	ID       int
	Email    string
	Password string
}

func Select() {
	var user User

	// 待执行语句
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// QueryRow(query string, args ...any) 函数接收两个参数:1.查询语句 2.查询参数列表(该参数是一个变长列表)
	// 返回Row，Row包含执行过程中的错误和返回的数据
	raw := db.QueryRow(userSql, 1)

	// Scan(dest ...any) 将匹配行中的列复制到 dest 指向的值中。
	// 如果有多行与查询匹配，Scan 将使用第一行并丢弃其余行。 如果没有行与查询匹配，Scan 将返回 ErrNoRows。简单理解就是将查询出的列对应到具体字段上。
	err := raw.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", user.Email, user.Password)
}

func Insert() {
	// 创建一个用户，设置要插入的数据
	var user User = User{ID: 4, Email: "110@qq.com", Password: "1234567890"}

	// 执行插入操作
	_, err := db.Exec("INSERT INTO users (id,email,password) VALUES($1,$2,$3)", user.ID, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	// 打印日志
	log.Printf("create ok!!!")

	// 测试数据是否插入成功，执行具体的查询语句
	var tmp User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// 设置查询参数为4，即创建数据时的ID值
	err = db.QueryRow(userSql, 4).Scan(&tmp.ID, &tmp.Email, &tmp.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// 输出查询结果
	fmt.Printf("Insert: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", tmp.Email, tmp.Password)
}

func Update() {
	// 创建一个用户，预要通过主键更改到数据库里
	var user User = User{ID: 4, Email: "dong@qq.com", Password: "abcdedf120"}

	// 执行更改操作
	_, err := db.Exec("UPDATE  users SET email=$1, password=$2 where id=$3", user.Email, user.Password, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	// 打印日志
	log.Printf("update ok!!!")

	// 测试数据是否更改成功，执行具体的查询语句
	var tmp User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// 设置查询参数为4，即把更改数据的ID
	err = db.QueryRow(userSql, 4).Scan(&tmp.ID, &tmp.Email, &tmp.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// 输出查询结果
	fmt.Printf("Update: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", tmp.Email, tmp.Password)
}

func Delete() {
	// 执行更改操作
	_, err := db.Exec("DELETE FROM  users  where id=$1", 4)
	if err != nil {
		log.Fatal(err)
	}

	// 打印日志
	log.Printf("delete ok!!!")

	// 测试数据是否删除成功，执行具体的查询语句
	var tmp User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// 设置查询参数为4，即被删除数据的ID
	err = db.QueryRow(userSql, 4).Scan(&tmp.ID, &tmp.Email, &tmp.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// 输出查询结果
	fmt.Printf("Delete: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", tmp.Email, tmp.Password)
}
