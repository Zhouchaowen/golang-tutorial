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
- 指针持有者类型：[Size]T、[]T、map[T]T、struct、func

```bigquery
bool
string

Numeric types:

uint        either 32 or 64 bits
int         same size as uint
uintptr     an unsigned integer large enough to store the uninterpreted bits of
            a pointer value
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers
            (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 (represents a Unicode code point)
```
## 运算符
Go支持五个基本二元算术运算符：+、-、*、/、%
Go支持六种位运算符：&、|、^、&^、<<、>>


## 参考
https://gfw.go101.org/article/operators.html

https://gfw.go101.org/article/type-system-overview.html



