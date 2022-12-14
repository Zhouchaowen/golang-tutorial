package main

import (
	"fmt"
)

// 占位符
func main() {
	var a byte = 255            // byte = uint8 rune = int32
	fmt.Printf("%v:%T\n", a, a) // 255:uint8

	var b int = 380
	// 不足位数前面补0
	fmt.Printf("%05d:%T\n", b, b)  // 00380:int
	fmt.Printf("%010d:%T\n", b, b) // 0000000380:int

	var c int = 88
	// 十进制 -> 二进制
	fmt.Printf("%b:%T\n", c, c) // 1011000:int
	// 十进制 -> 十六进制
	fmt.Printf("%x:%T\n", c, c) // 58:int

	var d string = "Golang"
	fmt.Printf("%s:%T\n", d, d) // Golang:string

	var e float64 = 3.14
	fmt.Printf("%f:%T\n", e, e) // 3.140000:float64
}

/*
[常用]
	%d    十进制表示
	%s    字符串或切片的无解译字节
	%f    有小数点而无指数，例如 123.456

	%v    相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
	%#v    相应值的 Go 语法表示
	%T    相应值的类型的 Go 语法表示
	%%    字面上的百分号，并非值的占位符

[布尔]

　　%t    单词 true 或 false。

[整数]

　　%b    二进制表示
　　%c    相应 Unicode 码点所表示的字符
　　%d    十进制表示
　　%o    八进制表示
　　%q    单引号围绕的字符字面值，由 Go 语法安全地转义
　　%x    十六进制表示，字母形式为小写 a-f
　　%X    十六进制表示，字母形式为大写 A-F
　　%U    Unicode 格式：U+1234，等同于 "U+%04X"

[浮点数及其复合构成]

　　%b    无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
　　%e    科学计数法，例如 -1234.456e+78
　　%E    科学计数法，例如 -1234.456E+78
　　%f    有小数点而无指数，例如 123.456
　　%g    根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
　　%G    根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出

[字符串与字节切片]

　　%s    字符串或切片的无解译字节
　　%q    双引号围绕的字符串，由 Go 语法安全地转义
　　%x    十六进制，小写字母，每字节两个字符
　　%X    十六进制，大写字母，每字节两个字符

[指针]

　　%p    十六进制表示，前缀 0x

[其它标记]

　　+    总打印数值的正负号；对于 %q（%+q）保证只输出 ASCII 编码的字符。
　　-    在右侧而非左侧填充空格（左对齐该区域）
　　#    备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或
　　0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始（即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的 Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
　　' '    （空格）为数值中省略的正负号留出空白（% d）；
                以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开：fmt.Printf("% x\n", "Hello")
// 参考 http://www.manongjc.com/detail/25-pmahqixdhaombky.html
*/
