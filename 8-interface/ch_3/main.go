package main

import (
	"fmt"
	"golang-tutorial/8-interface/ch_3/v1_struct"
	"golang-tutorial/8-interface/ch_3/v2_interface"
)

// 通过 _interface 构建一台 AMD CPU + 金士顿 Memory + 英特尔 NetWork + AOC Display 的电脑
func CreateComputer1() {
	cb := &v2_interface.ComputerBuilder{}
	cpu := v2_interface.AmdCPU{
		Name:       "Ryzen 5 5000",
		ModelType:  "十二线程",
		CoreNumber: 6,
	}
	mem := v2_interface.KingstonMemory{
		Name: "DDR4",
		Typ:  "金士顿",
		Cap:  16,
		MHz:  2666,
	}

	net := v2_interface.IntelNetWork{
		Name: "82574L",
		Typ:  "百兆以太网",
		Rate: 100,
	}

	dis := v2_interface.AOCDisplay{
		Name: "AOC",
		Typ:  "1080P",
	}
	c := cb.SetCPU(cpu).SetMemory(mem).SetNetWork(net).SetDisplay(dis).Build()
	c.RUN()
}

// 通过 _interface 构建一台 英特尔 CPU + 金士顿 Memory + 迈络思 NetWork + 飞利浦 Display 的电脑
func CreateComputer2() {
	cb := &v2_interface.ComputerBuilder{}
	cpu := v2_interface.IntelCPU{
		Name:       "i9-13900K",
		ModelType:  "二十四线程",
		CoreNumber: 12,
	}
	mem := v2_interface.KingstonMemory{
		Name: "DDR4",
		Typ:  "金士顿",
		Cap:  32,
		MHz:  2666,
	}

	net := v2_interface.MellanoxNetWork{
		Name: "82574L",
		Typ:  "千兆以太网",
		Rate: 1000,
	}

	dis := v2_interface.PhilipsDisplay{
		Name: "Philips",
		Typ:  "4K",
	}
	c := cb.SetCPU(cpu).SetMemory(mem).SetNetWork(net).SetDisplay(dis).Build()
	c.RUN()
}

// 通过 _struct 构建一台 AMD CPU + 金士顿 Memory + 英特尔 NetWork + AOC Display 的电脑
func CreateComputer3() {
	cb := &v1_struct.ComputerBuilder{}
	cpu := v1_struct.AmdCPU{
		Name:       "Ryzen 5 5000",
		ModelType:  "十二线程",
		CoreNumber: 6,
	}
	mem := v1_struct.KingstonMemory{
		Name: "DDR4",
		Typ:  "金士顿",
		Cap:  16,
		MHz:  2666,
	}

	net := v1_struct.IntelNetWork{
		Name: "82574L",
		Typ:  "百兆以太网",
		Rate: 100,
	}

	dis := v1_struct.AOCDisplay{
		Name: "AOC",
		Typ:  "1080P",
	}
	c := cb.SetCPU(cpu).SetMemory(mem).SetNetWork(net).SetDisplay(dis).Build()
	c.RUN()
}

// 构建一台 英特尔 CPU + 金士顿 Memory + 迈络思 NetWork + 飞利浦 Display 的电脑
func CreateComputer4() {
	cb := &v1_struct.ComputerBuilder2{}
	cpu := v1_struct.IntelCPU{
		Name:       "i9-13900K",
		ModelType:  "二十四线程",
		CoreNumber: 12,
	}
	mem := v1_struct.KingstonMemory{
		Name: "DDR4",
		Typ:  "金士顿",
		Cap:  32,
		MHz:  2666,
	}

	net := v1_struct.MellanoxNetWork{
		Name: "82574L",
		Typ:  "千兆以太网",
		Rate: 1000,
	}

	dis := v1_struct.PhilipsDisplay{
		Name: "Philips",
		Typ:  "4K",
	}
	c := cb.SetCPU(cpu).SetMemory(mem).SetNetWork(net).SetDisplay(dis).Build()
	c.RUN()
}

func main() {
	fmt.Println("------interface------")
	fmt.Println("低配: ")
	CreateComputer1()
	fmt.Println("高配: ")
	CreateComputer2()
	fmt.Println("------struct------")
	fmt.Println("低配: ")
	CreateComputer3()
	fmt.Println("高配: ")
	CreateComputer4()
}
