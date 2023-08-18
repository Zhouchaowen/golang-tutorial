package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

func init() {
	// Create DB pool
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
}

func main() {
	defer db.Close()

	ContextSelect()
}

func ContextSelect() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 如果上下文被取消或超时，查询执行将停止。
	_, err = db.QueryContext(ctx, "SELECT pg_sleep(15)")
	if err != nil {
		log.Fatal("query context err: ", err)
	}
	//打印日志
	log.Printf("delete ok!!!")
}
