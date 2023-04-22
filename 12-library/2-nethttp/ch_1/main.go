package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	URL    string
	Method string
)

func init() {
	flag.StringVar(&URL, "u", "https://www.baidu.com", "request rul")
	flag.StringVar(&Method, "m", "GET", "request method")
}

func main() {
	flag.Parse()

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)
}
