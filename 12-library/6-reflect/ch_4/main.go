package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	jsonStr := `{"id":4,"email":"golang@golang.cn","password":"xxxxx"}`
	users := Users{}
	if err := UnmarshalJSONToStruct(jsonStr, &users); err != nil {
		fmt.Println("错误:", err)
		return
	}

	fmt.Printf("%#v\n", users)
}

func UnmarshalJSONToStruct(jsonStr string, target interface{}) error {
	jsonStr = strings.ReplaceAll(jsonStr, "\"", "")
	jsonStr = strings.ReplaceAll(jsonStr, "`", "")
	jsonStr = strings.ReplaceAll(jsonStr, "{", "")
	jsonStr = strings.ReplaceAll(jsonStr, "}", "")
	jsonSlice := strings.Split(jsonStr, ",")

	jsonMp := make(map[string]string)
	for i, _ := range jsonSlice {
		tmp := strings.Split(jsonSlice[i], ":")
		jsonMp[tmp[0]] = tmp[1]
	}
	// 确保目标是一个非空的指向结构体的指针
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return fmt.Errorf("target must be a non-null pointer to a structure")
	}

	// 解引用指针以获取结构体值
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("data is not a struct")
	}

	t := value.Type()

	// 使用反射遍历结构体的字段
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := t.Field(i).Tag.Get("json") // 获取字段的JSON标签
		if fieldName == "" {
			continue
		}

		//属性的值 type
		switch field.Kind() {
		case reflect.Int:
			v := jsonMp[fieldName]
			vv, _ := strconv.ParseInt(v, 10, 64)
			field.SetInt(vv)
		case reflect.String:
			v := jsonMp[fieldName]
			field.SetString(v)
		default:
			panic("unsupported type")
		}

	}

	return nil
}
