package main

import "fmt"

type ComputerBuilder struct {
	Computer
}

type Computer struct {
	CPU
	Memory
	NetWork
	Display
}

func (c *ComputerBuilder) SetCPU(cpu CPU) *ComputerBuilder {
	c.CPU = cpu
	return c
}

func (c *ComputerBuilder) SetMemory(mem Memory) *ComputerBuilder {
	c.Memory = mem
	return c
}

func (c *ComputerBuilder) SetNetWork(nt NetWork) *ComputerBuilder {
	c.NetWork = nt
	return c
}

func (c *ComputerBuilder) SetDisplay(dis Display) *ComputerBuilder {
	c.Display = dis
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return c.Computer
}

func (c Computer) RUN() {
	c.CPU.operation()
	c.Memory.InteractiveData()
	c.NetWork.InteractiveData()
	c.Display.Display()
	fmt.Println("computer running")
}
