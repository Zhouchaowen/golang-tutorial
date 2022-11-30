package main

import "fmt"

type NetWork struct {
	name string
	typ  string
	rate int
}

func (n NetWork) InteractiveData() {
	fmt.Printf("%s %s %d is interactive data\n", n.name, n.typ, n.rate)
}
