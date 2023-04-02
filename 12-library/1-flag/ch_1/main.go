package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [args]\n", os.Args[0])
	flag.PrintDefaults()
}

// go run main.go -name zcw -debug -age 20 -score 80 -timeout 10m
func main() {
	// 定义一个字符串类型的命令行标志参数
	var name = flag.String("name", "defaultName", "the name to be used")
	// 定义一个布尔类型的命令行标志参数
	var debug = flag.Bool("debug", false, "enable debug mode")
	var age = flag.Int("age", 18, "the age of the person")
	var score = flag.Float64("score", 99.99, "the value of score")
	var timeout = flag.Duration("timeout", 5*time.Second, "the timeout duration")

	flag.Usage = usage

	// 解析命令行标志参数
	flag.Parse()
	fmt.Println("name:", *name, "debug:", *debug, "age:", *age, "score:", *score, "timeout:", *timeout)
}
