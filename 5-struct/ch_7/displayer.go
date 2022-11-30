package main

import "fmt"

type Display struct {
	name string
	typ  string
}

func (d Display) Display() {
	fmt.Printf("%s %s is display data\n", d.name, d.typ)
}
