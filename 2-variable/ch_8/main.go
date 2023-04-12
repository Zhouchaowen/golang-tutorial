package main

import "fmt"

func Steps1() {
	x := 10
	y := 3
	a := x + y // 13
	fmt.Printf("\t%d \n", a)
	a = x - y // 7
	fmt.Printf("\t%d \n", a)
	a = x * y // 30
	fmt.Printf("\t%d \n", a)
	a = x / y // 3
	fmt.Printf("\t%d \n", a)
	a = x % y // 1
	fmt.Printf("\t%d \n", a)
}

func Steps2() {
	x := 10
	y := 3
	b := x == y // false
	fmt.Printf("\t%t \n", b)
	b = x != y // true
	fmt.Printf("\t%t \n", b)
	b = x < y // false
	fmt.Printf("\t%t \n", b)
	b = x > y // true
	fmt.Printf("\t%t \n", b)
	b = x <= y // false
	fmt.Printf("\t%t \n", b)
	b = x >= y // true
	fmt.Printf("\t%t \n", b)
}

func Steps3() {
	x := true
	y := false
	c := x && y // false
	fmt.Printf("\t%t \n", c)
	c = x || y // true
	fmt.Printf("\t%t \n", c)
	c = !x // false
	fmt.Printf("\t%t \n", c)
}

func Steps4() {
	x := 0b1010 // 十进制数 10
	fmt.Printf("\t%05b %d\n", x, x)
	y := 0b0011 // 十进制数 3
	fmt.Printf("\t%05b %d\n", y, y)
	d := x & y // 0b0010，十进制数 2
	fmt.Printf("\t%05b %d\n", d, d)
	d = x | y // 0b1011，十进制数 11
	fmt.Printf("\t%05b %d\n", d, d)
	d = x ^ y // 0b1001，十进制数 9
	fmt.Printf("\t%05b %d\n", d, d)
	d = x << 1 // 0b10100，十进制数 20
	fmt.Printf("\t%05b %d\n", d, d)
	d = x >> 1 // 0b0101，十进制数 5
	fmt.Printf("\t%05b %d\n", d, d)
}

func Steps5() {
	x := 10
	x += 5 // x = 15
	fmt.Printf("\t%d \n", x)
	x -= 3 // x = 12
	fmt.Printf("\t%d \n", x)
	x *= 2 // x = 24
	fmt.Printf("\t%d \n", x)
	x /= 3 // x = 8
	fmt.Printf("\t%d \n", x)
	x %= 5 // x = 3
	fmt.Printf("\t%d \n", x)
	x &= 0b101 // x = 00001，十进制数 1
	fmt.Printf("\t%05b %d\n", x, x)
	x |= 0b110 // x = 00111，十进制数 7
	fmt.Printf("\t%05b %d\n", x, x)
	x ^= 0b011 // x = 00100，十进制数 4
	fmt.Printf("\t%05b %d\n", x, x)
	x <<= 1 // x = 01000，十进制数 8
	fmt.Printf("\t%05b %d\n", x, x)
	x >>= 2 // x = 00010，十进制数 2
	fmt.Printf("\t%05b %d\n", x, x)
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
	fmt.Println("Steps5():")
	Steps5()
}
