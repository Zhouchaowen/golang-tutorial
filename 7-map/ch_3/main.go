package main

import "fmt"

/*
	1.map深浅拷贝
*/

// Steps3 浅拷贝
func Steps3() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Println("\t-------------------------")
	tmpIntString := make(map[int]string, 2)
	tmpIntString = mpIntString
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

	tmpIntString[2] = "IMianBa"
	fmt.Println("\t-------------------------")
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
}

// Steps4 深拷贝
func Steps4() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Println("\t-------------------------")

	tmpIntString := make(map[int]string, 2)
	for k, v := range mpIntString {
		tmpIntString[k] = v
	}

	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

	tmpIntString[2] = "IMianBa"
	fmt.Println("\t-------------------------")
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
