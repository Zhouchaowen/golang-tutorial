# Map



## 目录

- ch_1 Map定义
- ch_2 初始化一个Map

## 定义映射

golang中通过map关键字定义映射, 格式`var mpIntInt map[T]T`表示定义了一个 Key 类型为 T，Value 类型为 T 的映射。

```go
package main

import "fmt"

// Steps1 定义映射
func Steps1() {
	// Steps 1-1: map[T]X 表示定义了一个 Key 类型为 T，Value 类型为 X 的映射
	// 定义一个 int->int 的map
	var mpIntInt map[int]int // 零值map
	fmt.Printf("\tmpIntInt:%+v len:%d\n",
		mpIntInt,
		len(mpIntInt)) // len 可以获取当前 map 存储的映射数量
	// mpIntInt[1] =1 // nil 映射不能添加键,添加报错 panic: assignment to entry in nil map

	// Steps 1-2: 定义一个 int->string 的map并初始化
	mpIntString := map[int]string{1: "Golang", 2: "Tutorial"}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 1-3: 用内建函数 make 来创建map
	mpIntBool := make(map[int]bool)
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))
	mpIntBool[0] = true
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))

	mpIntFloat32 := make(map[int]float32, 10)
	fmt.Printf("\tmpIntFloat32:%+v len:%d\n",
		mpIntFloat32,
		len(mpIntFloat32))

	mpStringSliceInt := make(map[string][]int, 10)
	fmt.Printf("\tmpStringSliceInt:%+v len:%d\n",
		mpStringSliceInt,
		len(mpStringSliceInt))
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
}
```

## 映射初始化与赋值

```go
package main

import "fmt"

// Steps2 map的基础使用
func Steps2() {
	// Steps 2-1: 用内建函数 make 来创建map
	mpIntString := make(map[int]string)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-2: 映射 mpIntString 中插入或修改元素
	mpIntString[0] = "Golang"
	mpIntString[1] = "World"
	mpIntString[1] = "Tutorial" // 修改mpIntString[1]元素
	mpIntString[2] = "Study"
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-3: 获取元素
	elem := mpIntString[0]
	fmt.Printf("\telem:%+v\n", elem)

	// Steps 2-4: 删除元素
	delete(mpIntString, 0)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-5: 通过双赋值检测某个键是否存在
	// 若 key 在 mpIntString 中，ok 为 true ; 否则, ok 为 false
	elem, ok := mpIntString[0]
	fmt.Printf("\telem:%+v ok:%t\n", elem, ok)

	// Steps 2-6: 通过range遍历map
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	for k, v := range mpIntString {
		fmt.Printf("\tKey:%d, Value:%s\n", k, v)
	}
}

func main() {
	fmt.Println("Steps2():")
	Steps2()
}
```


## 思考题
1. 通过Map实现一个人员统计小程序 
  
    1.1. 实现添加一个名字
   
    1.2. 实现查询一个名字是否存在

    1.3. 实现删除一个名字

    1.4. 实现更新一个名字
    
    1.5. 实现打印所有名字

    1.6. 实现统计总人数
    
2. 通过map,struct,func实现求一个班级所有学生最高总分,最低总分,各学科最高,最低分,平均分,打印所有学生所有学科成绩

```go
type Student struct {
	language float32
	math     float32
	english  float32
}

type class struct {
	mp map[string]Student
}

func (c class)ClassMaxScore() float64 {

	return 0
}

func (c class)ClassLanguageMaxScore() float64 {

	return 0
}

func (c class)PrintAllStudent()  {
	
}

func .....
```



## 参考
https://www.cnblogs.com/qcrao-2018/p/10903807.html#%E5%A6%82%E4%BD%95%E8%BF%9B%E8%A1%8C%E6%89%A9%E5%AE%B9

https://jonny.website/posts/go-map-bmap/

https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/#332-%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84
