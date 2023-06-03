package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
	1.定义全局结构体
	2.结构体赋值
	3.定义局部结构体
*/

// Demo 定义结构体
type Demo struct {
	// 小写表示不导出,包外不能引用
	a bool
	// 大写表示导出，包外能引用
	B byte
	C int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	D float32 // float64
	E string
	F []int
	G map[string]int
	H *int64
}

func Steps1() {
	d := Demo{ // 创建一个 Demo 类型的结构体实例
		a: true,
		B: 'b',
		C: 1,
		D: 1.0,
		E: "E",
		F: []int{1},
		G: map[string]int{"GOLANG": 1},
	}

	fmt.Printf("\td value %+v\n", d)

	fmt.Printf("\tdome.B: %c\n", d.B)

	// 访问结构体内的成员使用点. , 格式为：结构体变量.成员
	d.a = false // 修改a字段的值

	fmt.Printf("\td value %+v\n", d)
}

func Steps2() {
	// 结构体也可以定义在函数内
	type Demo struct {
		a int
		B string
	}

	d := Demo{ // 创建一个 Demo 类型的结构体实例
		a: 1,
	}

	fmt.Printf("\td value %+v\n", d)

	// 结构体字段使用点号来访问
	d.a = 2 // 修改a字段的值

	fmt.Printf("\td value %+v\n", d)
}

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func Steps3() {
	u := User{ // 创建一个 User 类型的结构体实例
		UserName: "golang",
		PassWord: "tutorial",
	}

	fmt.Printf("\tu value %+v\n", u)

	bytes, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("\tjson.Marshal error %s\n", err.Error())
	}
	fmt.Printf("\tjson user %s\n", string(bytes))
}

func Steps4() {
	u := User{ // 创建一个 User 类型的结构体实例
		UserName: "golang",
		PassWord: "tutorial",
	}
	t := reflect.TypeOf(u)              // 反射获取u的类型
	for i := 0; i < t.NumField(); i++ { // 通过类型获取结构体字段索引
		field := t.Field(i)
		fmt.Printf("\tfield %d: name=%s, json=%s \n", i, field.Name, field.Tag.Get("json"))
	}
}

func Steps5() {
	d := Demo{ // 创建一个 Demo 类型的结构体实例
		a: true,
		B: 'b',
		C: 1,
		D: 1.0,
		E: "E",
		F: []int{1},
		G: map[string]int{"GOLANG": 1},
	}

	// 结构体的字段内存地址排列
	fmt.Printf("\tvariable b   addr %p\n", &d)
	fmt.Printf("\tvariable b.a addr %p\n", &d.a)
	fmt.Printf("\tvariable b.B addr %p\n", &d.B)
	fmt.Printf("\tvariable b.C addr %p\n", &d.C)
	fmt.Printf("\tvariable b.D addr %p\n", &d.D)
	fmt.Printf("\tvariable b.E addr %p\n", &d.E)
	fmt.Printf("\tvariable b.F addr %p\n", &d.F)
	fmt.Printf("\tvariable b.G addr %p\n", &d.G)

	fmt.Printf("\t-----------------\n")

	c := d
	fmt.Printf("\tvariable c   addr %p\n", &c)
	fmt.Printf("\tvariable c.a addr %p\n", &c.a)
	fmt.Printf("\tvariable c.B addr %p\n", &c.B)
	fmt.Printf("\tvariable c.C addr %p\n", &c.C)
	fmt.Printf("\tvariable c.D addr %p\n", &c.D)
	fmt.Printf("\tvariable c.E addr %p\n", &c.E)
	fmt.Printf("\tvariable c.F addr %p\n", &c.F)
	fmt.Printf("\tvariable c.G addr %p\n", &c.G)
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
