package main

import "fmt"

type CPU struct {
	name       string
	modelType  string
	coreNumber int
}

func (c CPU) operation() {
	fmt.Printf("%s %s %d is operation\n", c.name, c.modelType, c.coreNumber)
}
