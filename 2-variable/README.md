# 目录
- ch_1 变量定义
- ch_2 变量赋值
- ch_3 类型转换
- ch_4 常量定义
- ch_5 定义函数变量
- ch_6 定义指针变量
- ch_7 占位符

## 数据类型
- 布尔类型：bool
- 整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。
- 浮点数类型：float32、float64。
- 复数类型：complex64、complex128。
- 字符串类型：string。
- 应用类型 [Size]T、[]T、map[T]T、struct、func

```bigquery
*T         // 一个指针类型
[5]T       // 一个元素类型为T、元素个数为5的数组类型
[]T        // 一个元素类型为T的切片类型
map[Tkey]T // 一个键值类型为Tkey、元素类型为T的映射类型

// 一个结构体类型
struct {
	name string
	age  int
}

// 一个函数类型
func(int) (bool, string)

// 一个接口类型
interface {
	Method0(string) int
	Method1() (int, bool)
}

// 几个通道类型
chan T
chan<- T
<-chan T
```
## 运算符
Go支持五个基本二元算术运算符：+、-、*、/、%
Go支持六种位运算符：&、|、^、&^、<<、>>


## 参考
https://gfw.go101.org/article/operators.html

https://gfw.go101.org/article/type-system-overview.html



