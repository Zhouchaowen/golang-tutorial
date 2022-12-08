package _struct

import "fmt"

// 构建一台 英特尔 CPU + 金士顿 Memory + 英特尔 NetWork + AOC Display 的电脑
type ComputerBuilder struct {
	Computer1
}

type Computer1 struct {
	AmdCPU
	KingstonMemory
	IntelNetWork
	AOCDisplay
}

func (c *ComputerBuilder) SetCPU(cpu AmdCPU) *ComputerBuilder {
	c.AmdCPU = cpu
	return c
}

func (c *ComputerBuilder) SetMemory(mem KingstonMemory) *ComputerBuilder {
	c.KingstonMemory = mem
	return c
}

func (c *ComputerBuilder) SetNetWork(nt IntelNetWork) *ComputerBuilder {
	c.IntelNetWork = nt
	return c
}

func (c *ComputerBuilder) SetDisplay(dis AOCDisplay) *ComputerBuilder {
	c.AOCDisplay = dis
	return c
}

func (c *ComputerBuilder) Build() Computer1 {
	return c.Computer1
}

func (c Computer1) RUN() {
	c.AmdCPU.operation()
	c.KingstonMemory.InteractiveData()
	c.IntelNetWork.TransferData()
	c.AOCDisplay.Display()
	fmt.Println("\tcomputer running")
}

// 构建一台 AMD CPU + 金士顿 Memory + 迈络思 NetWork + 飞利浦 Display 的电脑
type ComputerBuilder2 struct {
	Computer2
}

type Computer2 struct {
	IntelCPU
	KingstonMemory
	MellanoxNetWork
	PhilipsDisplay
}

func (c *ComputerBuilder2) SetCPU(cpu IntelCPU) *ComputerBuilder2 {
	c.IntelCPU = cpu
	return c
}

func (c *ComputerBuilder2) SetMemory(mem KingstonMemory) *ComputerBuilder2 {
	c.KingstonMemory = mem
	return c
}

func (c *ComputerBuilder2) SetNetWork(nt MellanoxNetWork) *ComputerBuilder2 {
	c.MellanoxNetWork = nt
	return c
}

func (c *ComputerBuilder2) SetDisplay(dis PhilipsDisplay) *ComputerBuilder2 {
	c.PhilipsDisplay = dis
	return c
}

func (c *ComputerBuilder2) Build() Computer2 {
	return c.Computer2
}

func (c Computer2) RUN() {
	c.IntelCPU.operation()
	c.KingstonMemory.InteractiveData()
	c.MellanoxNetWork.TransferData()
	c.PhilipsDisplay.Display()
	fmt.Println("\tcomputer running")
}
