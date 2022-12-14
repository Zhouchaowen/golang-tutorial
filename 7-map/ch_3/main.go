package main

import "fmt"

/**
1.map当做参数传递
*/

func addMap(mp map[int]string) {
	fmt.Printf("\tmp value addr:%p\n", mp)
	fmt.Printf("\tmp addr:%p\n", &mp)
	mp[0] = "0"
	mp[1] = "1"
	mp[2] = "2"
	mp[3] = "3"
	fmt.Printf("\tmp value addr:%p\n", mp)
	fmt.Printf("\tmp addr:%p\n", &mp)
}

// Steps3
func Steps3() {
	// Steps 2-1: 用内建函数 make 来创建map
	mpIntString := make(map[int]string, 2)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Printf("\tmpIntString addr:%p\n", &mpIntString)
	addMap(mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Printf("\tmpIntString addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
}
