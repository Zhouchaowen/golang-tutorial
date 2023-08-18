package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // 注册psql数据库驱动
	"log"
	"reflect"
	"strings"
	"time"
)

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.105:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

type Users struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func init() {
	// 创建数据库连接
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
}

func main() {
	user := Users{
		ID:       4,
		Email:    "golang@golang.cn",
		Password: "xxxxx",
	}

	// 生成 insert sql
	insertSQL, args := GenerateInsertSQL(user, "")
	fmt.Println("INSERT SQL:", insertSQL)
	fmt.Println("ARGS:", args)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 执行插入操作
	_, err := db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		log.Fatal(err)
	}

	// 打印日志
	log.Printf("create ok!!!")

	// 测试数据是否插入成功，执行具体的查询语句
	var tmp Users
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// 设置查询参数为4，即创建数据时的ID值
	err = db.QueryRow(userSql, 4).Scan(&tmp.ID, &tmp.Email, &tmp.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	// 输出查询结果
	fmt.Printf("Insert: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", tmp.Email, tmp.Password)
}

func GenerateInsertSQL(data interface{}, tableName string) (string, []interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		panic("GenerateInsertSQL: data is not a struct")
	}

	t := value.Type()

	if tableName == "" {
		tableName = strings.ToLower(t.Name())
	}

	var columns []string
	var placeholders []string
	var args []interface{}

	var idx = 1
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := t.Field(i).Tag.Get("db")
		if tag == "" {
			continue
		}

		// 判断是否可以使用 Interface 而不会出现恐慌
		if field.CanInterface() {
			columns = append(columns, tag)
			placeholders = append(placeholders, fmt.Sprintf("$%d", idx))
			args = append(args, field.Interface())

			idx++
		}
	}

	columnsStr := strings.Join(columns, ",")
	placeholdersStr := strings.Join(placeholders, ",")

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, columnsStr, placeholdersStr)
	return sql, args
}
