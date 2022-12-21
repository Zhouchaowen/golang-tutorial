package main

import (
	"flag"
	"fmt"
)

var (
	Url      string
	FileName string
)

func init() {
	flag.StringVar(&Url, "u", "https://www.baidu.com", "url path")
	flag.StringVar(&FileName, "f", "test.png", "save png file name")
}

func main() {
	flag.Parse()

	fmt.Println(Url)
	fmt.Println(FileName)
}
