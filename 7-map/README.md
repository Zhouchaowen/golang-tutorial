# Map

`map`是`Go`语言中的一种内置数据结构，也称为**哈希表或字典**。它是一种**无序的键值对集合**，其中每个键唯一对应一个值，通过键来快速查找对应的值。在`map`中，所有的键都必须是同一类型，所有的值也必须是同一类型。

## 目录

- 定义 Map与赋值

## 定义 Map与赋值

`map`的声明方式：

```go
var mapName map[keyType]valueType
```

其中`mapName`表示变量名称，`keyType`表示键的类型，`valueType`表示值的类型。需要注意的是，这里只是声明了`map`变量并没有初始化，如果直接使用将会引发运行时错误。正确的初始化方式是使用`make`函数：

```go
mapName := make(map[keyType]valueType,[size]) // size 可选
```

或者直接定义或初始化默认值：

```go
mapName := map[keyType]valueType{}
mapName := map[keyType]valueType{
    key1: value1,
    key2: value2,
    ...
}
```

在map中添加或修改键值对可以直接使用下标运算符`[]`，例如：

```go
mapName[key] = value
```

如果键值对中的键已经存在，那么它所对应的值就会被更新；如果不存在，就会添加一个新的键值对。

查找map中的键值对也可以使用下标运算符`[]`，例如：

```go
value := mapName[key]
```

这个操作会返回键所对应的值，如果该键不存在，会返回该值类型的零值。

具体示例：

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
	// mpIntInt[1] =1 // 未初始化的映射不能添加键,添加报错 panic: assignment to entry in nil map

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

`map`还有一些其他的函数：

- `delete(mapName, key)`：删除指定键的键值对。
- `len(mapName)`：返回map中键值对的数量。
- `for key, value := range mapName`：遍历map中的键值对。

需要注意的是，由于`map`是无序的，所以遍历时不能保证顺序。

具体示例：

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

## 自检

- `map`的定义和声明 ？
- `map`的初始化 ？
- `map`的元素访问和赋值 ？
- `map`的遍历 ？
- `map`的长度和删除元素 ？
- `map`的传递方式 ?

## 参考

https://www.cnblogs.com/qcrao-2018/p/10903807.html

https://jonny.website/posts/go-map-bmap/

https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap
