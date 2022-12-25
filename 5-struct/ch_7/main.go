package main

/*
	1.结构体组合
*/

func main() {
	cb := &ComputerBuilder{}
	cpu := CPU{
		name:       "AMD Ryzen 5 5000",
		modelType:  "十二线程",
		coreNumber: 6,
	}
	mem := Memory{
		name: "DDR4",
		typ:  "金百达",
		cap:  32,
		mHz:  2666,
	}

	net := NetWork{
		name: "Intel 82574L",
		typ:  "千兆以太网",
		rate: 1000,
	}

	dis := Display{
		name: "AOC",
		typ:  "4K",
	}
	c := cb.SetCPU(cpu).SetMemory(mem).SetNetWork(net).SetDisplay(dis).Build()
	c.RUN()
}
