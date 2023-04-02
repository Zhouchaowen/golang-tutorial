package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	name    string
	debug   bool
	age     int
	score   float64
	timeout time.Duration
)

// go run main.go -name zcw -debug -age 20 -score 80 -timeout 10m
func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&name, "name", "defaultName", "the name to be used")
	flag.IntVar(&age, "age", 18, "the age of the person")
	flag.Float64Var(&score, "score", 99.99, "the value of score")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "the timeout duration")
}

func main() {
	// 解析命令行标志参数
	flag.Parse()
	fmt.Println("name:", name, "debug:", debug, "age:", age, "score:", score, "timeout:", timeout)
}
