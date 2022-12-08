package _struct

import "fmt"

type IntelCPU struct {
	Name       string
	ModelType  string
	CoreNumber int
}

func (c IntelCPU) operation() {
	fmt.Printf("\tIntel %s %s %d is operation\n", c.Name, c.ModelType, c.CoreNumber)
}

type AmdCPU struct {
	Name       string
	ModelType  string
	CoreNumber int
}

func (c AmdCPU) operation() {
	fmt.Printf("\tAMD %s %s %d is operation\n", c.Name, c.ModelType, c.CoreNumber)
}
