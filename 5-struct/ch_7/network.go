package main

import "fmt"

type NetWork struct {
	name string
	typ  string
	rate int
}

func (n NetWork) TransferData() {
	fmt.Printf("%s %s %d is transfer data\n", n.name, n.typ, n.rate)
}
