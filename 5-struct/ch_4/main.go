package main

import "fmt"

// ResponseStatus 自定义类型的方法
type ResponseStatus int

const (
	QuerySuccess ResponseStatus = iota
	QueryError
)

func (r ResponseStatus) ToCN() string {
	switch r {
	case 0:
		return "query success"
	case 1:
		return "query error"
	default:
		return "non"
	}
}

func main() {
	fmt.Println(QuerySuccess.ToCN())
	fmt.Println(QueryError.ToCN())
}
