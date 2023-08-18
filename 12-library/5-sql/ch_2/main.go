package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
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

	PrepareSelect()
	PrepareDelete()
	PrepareInsert()
	PrepareBenchmark()
}

func PrepareSelect() {
	var user User

	// Prepare(query string) (*Stmt, error) Prepare函数为以后要查询或执行的操作预先准备好sql。
	// 多个查询或执行可以通过返回的Stmt并发执行。 当不再需要该语句时，必须调用该语句的 Close 方法。
	stmt, err := db.Prepare("SELECT id, email, password FROM users WHERE id = $1")
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
	}
	defer stmt.Close() // 准备好的语句会占用服务器资源，使用后应关闭

	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal("Failed to Prepare query: ", err)
	}

	// 查询返回的列
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("rows.Columns error: %v\n", err)
	}
	fmt.Printf("columns: %v\n", columns)

	// 查询返回列对应属性
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatalf("rows.ColumnTypes error: %v\n", err)
	}
	for _, v := range columnTypes {
		fmt.Printf("types: %v\n", v)
	}

	// 获取返回查询列的值
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			log.Fatalf("rows.Scan error: %v\n", err)
		}
		fmt.Printf("Select: 你好 邮箱：%s, 密码：%s,  欢迎回来!\n", user.Email, user.Password)
	}
}

func PrepareInsert() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	stmt, err := db.Prepare("INSERT INTO users (id,email,password) VALUES($1,$2,$3)")
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

}

func PrepareDelete() {
	users := []User{
		{ID: 10, Email: "1101@qq.com", Password: "1234567890"},
		{ID: 11, Email: "1102@qq.com", Password: "1234567890"},
		{ID: 12, Email: "1103@qq.com", Password: "1234567890"},
		{ID: 13, Email: "1104@qq.com", Password: "1234567890"},
	}

	stmt, err := db.Prepare("DELETE FROM users  where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, user := range users {
		if result, err := stmt.Exec(user.ID); err != nil {
			log.Fatal(err)
		} else {
			rowsAffected, _ := result.RowsAffected()
			lastInsertId, _ := result.LastInsertId()
			fmt.Printf("rowsAffected:%d, lastInsertId：%d\n", rowsAffected, lastInsertId)
		}
	}

}

// PrepareBenchmark 重复利用 Prepare 后的 stmt 语句，可以有效提高 SQL 执行性能
func PrepareBenchmark() {
	var num int
	start := time.Now()
	for i := 0; i < 100; i++ {
		err = db.QueryRow("select count(*) from pg_stat_activity where datname = $1", "postgres").Scan(&num)
		if err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("got: %d in %s\n", num, elapsed)

	stmt, err := db.Prepare("select count(*) from pg_stat_activity where datname = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	start = time.Now()
	for i := 0; i < 100; i++ {
		err = stmt.QueryRow("postgres").Scan(&num)
	}
	elapsed = time.Since(start)
	fmt.Printf("got: %d in %s\n", num, elapsed)
}
