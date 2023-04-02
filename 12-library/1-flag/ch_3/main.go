package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
)

var (
	name    string
	debug   bool
	age     int
	score   float64
	timeout time.Duration
)

// go run main.go --name zcw --debug -a 20 -s 80 -t 10m
// 短命令和长命令需--和-区分
func init() {
	pflag.BoolVar(&debug, "debug", false, "enable debug mode")
	pflag.StringVar(&name, "name", "defaultName", "the name to be used")

	pflag.IntVarP(&age, "age", "a", 18, "the age of the person")
	pflag.Float64VarP(&score, "score", "s", 99.99, "the value of score")
	pflag.DurationVarP(&timeout, "timeout", "t", 5*time.Second, "the timeout duration")
}

func main() {
	// 解析命令行标志参数
	pflag.Parse()
	fmt.Println("name:", name, "debug:", debug, "age:", age, "score:", score, "timeout:", timeout)
}
