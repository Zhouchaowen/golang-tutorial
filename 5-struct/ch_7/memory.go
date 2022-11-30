package main

import "fmt"

type Memory struct {
	name string
	typ  string
	cap  int
	mHz  int
}

func (m Memory) InteractiveData() {
	fmt.Printf("%s %s %d %d is interactive data\n", m.name, m.typ, m.cap, m.mHz)
}
