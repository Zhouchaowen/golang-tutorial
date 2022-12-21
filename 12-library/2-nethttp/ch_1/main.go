package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	Ip   string
	Port int
)

func init() {
	flag.StringVar(&Ip, "i", "0.0.0.0", "server ip")
	flag.IntVar(&Port, "p", 8080, "server port")
}

func main() {
	flag.Parse()
	mu := http.NewServeMux()
	mu.Handle("/hello", printHandler())

	mu.HandleFunc("/world", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "world")
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", Ip, Port), mu)
}

func printHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
}
