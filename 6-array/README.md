# 数组与切片

## 目录

- ch_1 数组基础用法
- ch_2 切片基础用法



## 数组基础用法

golang中定义数组必须指定大小如: `var arrayInt = [3]int{}`定义大小为3的`int`型数组

```go
package main

import "fmt"

type dome struct {
	a int
	b float32
}

// 定义数组, 数组必须指定大小
func main() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	// 类型 [3]int 表示拥有 3 个 int 类型的值的数组, 默认值为0
	var arrayInt = [3]int{} // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	arrayInt[0] = 1
	arrayInt[1] = 2
	fmt.Printf("arrayInt: %+v\n", arrayInt)

	arrayBool := [3]bool{false, true}
	fmt.Printf("arrayBool: %+v\n", arrayBool)

	arrayFloat32 := [3]float32{1.0, 2.0} // float64
	fmt.Printf("arrayFloat32: %+v\n", arrayFloat32)

	arrayString := [3]string{"Golang", "Tutorial"}
	fmt.Printf("arrayString: %+v\n", arrayString)

	arrayStruct := [3]dome{{a: 1, b: 2.0}, {a: 11, b: 22.0}}
	fmt.Printf("arrayStruct: %+v\n", arrayStruct)

	// 数组可以直接通过下标访问 T[x]
	fmt.Printf("arrayInt[0]: %d\n", arrayInt[0])

	// 数组可以直接通过下标修改 T[x] = y
	arrayInt[0] = 11
	fmt.Printf("arrayInt[0]: %d\n", arrayInt[0])
}
```

## 切片基础用法

### 定义切片

```go
package main

import (
	"fmt"
	"unsafe"
)

// 切片也可以定义在全局
var sliceByte []byte

// Steps1 定义切片
func Steps1() {
	// Steps 1-1: 类型 []T 表示一个元素类型为 T 的切片
	// 切片拥有长度和容量, 切片在添加数据时会自动扩容, 可以通过len(),cap()获取切片长度和容量
	var sliceInt []int // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr

	// Steps 1-2: append 向切片中添加元素（可能会导致内存重新分配）
	for i := 0; i < 10; i++ {
		sliceInt = append(sliceInt, i)
	}
	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	// Steps 1-3: 获取切片长度
	fmt.Println("\tsliceInt len:", len(sliceInt))

	// Steps 1-4: 获取切片的容量
	fmt.Println("\tsliceInt cap:", cap(sliceInt))

	// Steps 1-5: nil 切片的长度和容量为 0 且没有底层数组
	var sliceBool []bool
	fmt.Printf("\tsliceBool:%+v len:%d cap:%d\n",
		sliceBool,
		len(sliceBool),
		cap(sliceBool))
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	fmt.Println("Steps1():")
	Steps1()
}
```

### 初始化切片

```go
package main

import (
	"fmt"
	"unsafe"
)

// Steps2 定义并初始化切片
func Steps2() {
	// Steps 2-1: 初始化切片
	sliceString := []string{"Golang", "Tutorial"}
	fmt.Printf("\tsliceString:%+v len:%d cap:%d\n",
		sliceString,
		len(sliceString),
		cap(sliceString))
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	fmt.Println("Steps2():")
	Steps2()
}
```

### 通过make创建切片

```go
package main

import (
	"fmt"
	"unsafe"
)

// Steps3 通过 make 创建切片
func Steps3() {
	// Steps 3-1: 用内建函数 make 来创建切片
	// make([]T,len,cap)
	sliceFloat32 := make([]float32, 5)
	fmt.Printf("\tsliceFloat32:%+v len:%d cap:%d\n",
		sliceFloat32,
		len(sliceFloat32),
		cap(sliceFloat32))
	sliceFloat64 := make([]float64, 5, 10)
	fmt.Printf("\tsliceFloat64:%+v len:%d cap:%d\n",
		sliceFloat64,
		len(sliceFloat64),
		cap(sliceFloat64))
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	Steps3()
	fmt.Println("Steps4():")
}
```

### 定义二维切片

```go
package main

import (
	"fmt"
	"unsafe"
)

// Steps4 二维切片
func Steps4() {
	// Steps 4-1: 定义二维切片，并赋值
	sliceStringString := [][]string{
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
	}
	fmt.Printf("\tsliceStringString:%+v len:%d cap:%d\n",
		sliceStringString,
		len(sliceStringString),
		cap(sliceStringString))
	// Steps 4-3: 添加一行
	sliceStringString = append(sliceStringString, []string{"1", "1", "1", "1", "1"})
	fmt.Printf("\tsliceStringString:%+v len:%d cap:%d\n",
		sliceStringString,
		len(sliceStringString),
		cap(sliceStringString))

	// Steps 4-3: 打印二维数组
	for i := 0; i < len(sliceStringString); i++ { // len(sliceStringString) y轴数组长度
		fmt.Print("\t")
		for j := 0; j < len(sliceStringString[i]); j++ { // len(sliceStringString[i]) 第i行 x轴数组长度
			fmt.Printf("%s ", sliceStringString[i][j])
		}
		fmt.Println()
	}
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	fmt.Println("Steps4():")
	Steps4()
}
```

### 截取切片

```go
package main

import (
	"fmt"
	"unsafe"
)

// Steps5 切片上截取切片
func Steps5() {
	// Steps 5-1: 定义切片并初始化
	sliceInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	// Steps 5-2: 可以用 slice[low : high] or slice[low : high] 来截取数组或切片的一个片段长度为 high-low
	// 注意: sliceInt[0:3] 等同于 sliceInt[:3]
	interceptionSliceInt := sliceInt[1:3] // 获取 sliceInt 下标 1-2 的元素:[1,2,3] 长度为2
	fmt.Printf("\tinterceptionSliceInt:%+v len:%d cap:%d\n",
		interceptionSliceInt,
		len(interceptionSliceInt),
		cap(interceptionSliceInt))

	// Steps 5-3: 可以用 slice[low : high: cap] 来截取切片或数组的一个片段长度为 high-low,容量为cap
	interceptionSliceIntCap := sliceInt[1:3:5] // 获取 sliceInt 下标 1-2 的元素:[1,2,3] 长度为2, 容量为4
	fmt.Printf("\tinterceptionSliceIntCap:%+v len:%d cap:%d\n",
		interceptionSliceIntCap,
		len(interceptionSliceIntCap),
		cap(interceptionSliceIntCap))

	// Steps 5-4: 切片并不存储任何数据，它只是描述了底层数组中的一段
	// 更改切片的元素会修改其底层数组中对应的元素,与它共享底层数组的切片都会观测到这些修改

	// interceptionSliceIntCap[2] 超出当前len, 打印报错 panic: runtime error: index out of range [2] with length 2
	//fmt.Printf("interceptionSliceIntCap[2]:%d",interceptionSliceIntCap[2])

	// 通过指针偏移强行获取底层元素（这种方式时不安全的）
	fmt.Printf("\tinterceptionSliceCap[2]:%d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionSliceIntCap[0])) + uintptr(16))))

	// Steps 5-6: 修改interceptionSliceCap[2]的值为33,底层切片sliceInt对应[3]位置改变33
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionSliceIntCap[0])) + uintptr(16))) = 33
	fmt.Printf("\tsliceInt[3]:%d\n", sliceInt[3])

	interceptionSliceIntCap[0] = 11
	fmt.Printf("\tsliceInt[1]:%d\n", sliceInt[1])
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	fmt.Println("Steps5():")
	Steps5()
}
```

## 数组的拷贝

```go
package main

import (
	"fmt"
)

// 指针持有者类型的拷贝问题

// Steps1 浅拷贝
func Steps1() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sliceIntTmp []int

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	sliceIntTmp = sliceInt
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))

	sliceIntTmp[0] = 111

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))
}

// Steps2 深拷贝
func Steps2() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sliceIntTmp []int

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	sliceIntTmp = make([]int, len(sliceInt))

	copy(sliceIntTmp, sliceInt) // 深拷贝

	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))

	sliceIntTmp[0] = 111

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
}
```

## 数组与切片参数传递时的区别

```go
package main

import (
	"fmt"
)

func modifySlice0(arr []int) {
	arr[0] = 1000
}

// 切片作为函数参数时传递的是指针类型的值
func Steps3() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	modifySlice0(sliceInt)

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
}

func modifyArr0(arr [10]int) {
	arr[0] = 1000
}

// 数组作为函数参数时传递的是值类型的全拷贝
func Steps4() {
	var sliceInt = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	modifyArr0(sliceInt)

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
```



## 思考题

1. 定义一个方法求出数组中奇数和偶数的和, 并同时返回。
2. 定义一个int型大小为5的自定义类型数组, 并定义打印所有元素的方法和求和方法。
```go
type myInt []int
```

3. 计算任意两个20位的整数的加减乘除

```go
12345678912345678912+12345678912345678912
```

4. 通过slice,struct,func实现求一个班级所有学生最高总分,最低总分,各学科最高,最低分,平均分

```go
type Student struct {
  name     string
	language float32
	math     float32
	english  float32
}

type class struct {
	students []Student
}

func ClassMaxScore(students []Student) float64 {

	return 0
}

func ClassLanguageMaxScore(students []Student) float64 {

	return 0
}

func .....
```

## 参考

https://gfw.go101.org/article/value-part.html

https://tour.go-zh.org/moretypes/7

https://blog.go-zh.org/go-slices-usage-and-internals

https://emmie.work/posts/golang-slice-assignment-%E8%88%87-append-%E6%96%B9%E6%B3%95/

https://ueokande.github.io/go-slice-tricks/

https://divan.dev/posts/avoid_gotchas/

https://juejin.cn/post/7055660145988075550

https://www.practical-go-lessons.com/chap-21-slices