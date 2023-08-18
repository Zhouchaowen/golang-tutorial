package main

import (
	"fmt"
	"reflect"
)

/*
	1.介绍reflect.TypeOf(u)
	2.Kind()
	3.NumField()
	4.StructField
*/

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (u User) PrintSex(sex string) {
	fmt.Println("sex:", sex)
}

func (u User) printAge(age string) {
	fmt.Println("age:", age)
}

func Steps1() {
	u := User{ // 创建一个 User 类型的结构体实例
		UserName: "golang",
		PassWord: "tutorial",
	}

	// reflect.TypeOf(u) 返回 reflect.Type 表示u的类型
	t := reflect.TypeOf(u)

	// t.Kind() 返回 Kind 类型, 该类型枚举了 Go 中所有类型
	if t.Kind() != reflect.Struct {
		return
	}

	// t.NumField() 返回 struct 字段数量
	for i := 0; i < t.NumField(); i++ {
		// t.Field(i) 返回当前索引的字段信息
		field := t.Field(i)
		// 获取field类型,tag的值
		fmt.Printf("\tfield %d: name=%s, type=%s, json=%s \n", i, field.Name, field.Type.Kind(), field.Tag.Get("json"))
	}

	// t.FieldByName(x) 返回 struct 名称为 x 的字段
	field, ok := t.FieldByName("UserName")
	if ok {
		fmt.Printf("\tname=%s, type=%s, json=%s \n", field.Name, field.Type.Kind(), field.Tag.Get("json"))
	}
}

func Steps2() {
	u := User{ // 创建一个 User 类型的结构体实例
		UserName: "golang",
		PassWord: "tutorial",
	}

	// reflect.TypeOf(u) 返回 reflect.Type 表示u的类型
	t := reflect.TypeOf(u)

	// Kind() 返回具体类型
	fmt.Println("Kind:", t.Kind())

	// NumMethod() 返回使用方法可访问的方法数。对于非接口类型，它返回导出方法的数量; 对于接口类型，它返回导出和未导出方法的数量。
	fmt.Printf("NumMethod: %d\n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		// 获取到索引为 i 的方法
		m := t.Method(i)

		fmt.Printf("Name:%s\n", m.Name)               // 方法名称
		fmt.Printf("IsExported:%t\n", m.IsExported()) // 方法是否为导出

		typ := m.Type
		fmt.Printf("Type:%s\n", typ.Kind()) // 方法是否为导出

		fuc := m.Func
		fuc.Call([]reflect.Value{reflect.ValueOf(u), reflect.ValueOf("man")})
	}
}

func Steps3() {
	// 创建一个 User 类型的结构体指针
	u := User{
		UserName: "golang",
		PassWord: "tutorial",
	}

	// reflect.TypeOf(u) 返回 reflect.Type 表示u的类型
	t := reflect.ValueOf(u)

	// Kind() 返回具体类型
	fmt.Println("Kind:", t.Kind())

	// t.NumField() 返回 struct 字段数量
	for i := 0; i < t.NumField(); i++ {
		// t.Field(i) 返回当前索引的字段信息
		field := t.Field(i)

		// 获取field值，类型
		fmt.Printf("\tfield %d: value=%s, kind=%s \n", i, field.Interface(), field.Kind())
	}

	// 获取字段名称为 UserName 的字段与reflect.TypeOf(u)的有所不同
	v := t.FieldByName("UserName")
	fmt.Println("Filed UserName Kind:", v.Kind())

	fmt.Println("UserName:", v.Interface())
}

func Steps5() {
	u := User{ // 创建一个 User 类型的结构体实例
		UserName: "golang",
		PassWord: "tutorial",
	}

	// reflect.TypeOf(u) 返回 reflect.Type 表示u的类型
	t := reflect.TypeOf(u)

	// Kind() 返回具体类型
	fmt.Println("Kind: ", t.Kind())

	// Align() 返回在内存中分配此类型值时的对齐方式（以字节为单位）
	fmt.Printf("Align: %d byte \n", t.Align())

	// FieldAlign() 当用作结构体中的字段时，返回此类型值的字节对齐方式。(把t类型当做其它结构体的字段时的对齐字节)
	fmt.Printf("FieldAlign: %d byte \n", t.FieldAlign())

	// NumMethod() 返回使用方法可访问的方法数。对于非接口类型，它返回导出方法的数量; 对于接口类型，它返回导出和未导出方法的数量。
	fmt.Printf("NumMethod: %d\n", t.NumMethod())

	// NumField() 返回结构类型的字段计数,如果类型不是struct，它会崩溃。
	fmt.Printf("NumField:  %d\n", t.NumField())

	// Name() 返回已定义类型的包中的类型名称。对于其他（未定义）类型，它返回空字符串。
	fmt.Printf("Name: %s\n", t.Name())

	// PkgPath() 返回已定义类型的包路径，即唯一标识包的导入路径。
	fmt.Printf("PkgPath: %s\n", t.PkgPath())

	// String() 返回包名/类型名
	fmt.Println("String: ", t.String())

	// Size() 返回存储给定类型的值所需的字节数
	fmt.Printf("Size: %d\n", t.Size())

	// Comparable() 此类型的值是否可比较
	fmt.Println("Comparable: ", t.Comparable())

	// Field(idx) 返回结构类型中索引是idx的字段,如果类型不是struct，它会崩溃
	fmt.Printf("Field: %+v\n", t.Field(0))

	// FieldByName(name) 返回具有给定名称的结构字段和一个指示是否找到该字段的布尔值
	structField, ok := t.FieldByName("UserName")
	fmt.Printf("FieldByName: %+v, ok: %t\n", structField, ok)
}

func main() {
	//Steps1()
	//Steps2()
	Steps3()
}
