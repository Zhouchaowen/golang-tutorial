package ch_1

func add(x, y int) int {
	return x + y
}

func memoryAllocation(x int) {
	_ = make([]int, x)
	return
}
