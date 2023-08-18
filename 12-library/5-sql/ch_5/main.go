package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const (
	DB_DSN = "postgres://postgres:lab_password@10.2.0.104:5432/postgres?sslmode=disable"
)

var db *sql.DB
var err error

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// 将请求上下文传递给slowQuery()，以便它可以用作父上下文。
	err := slowQuery(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}

	fmt.Fprintln(w, "OK")
}

func slowQuery(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "SELECT pg_sleep(10)")
	return err
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %s", err.Error())
	http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
}

func main() {
	db, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", exampleHandler)

	log.Print("Listening...")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal(err)
	}

	// 启动程序后执行，检验查询超时 curl -i localhost:5000/
}
