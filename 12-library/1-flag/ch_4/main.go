package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("12-library/1-flag/ch_4/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件失败的情况
		log.Panicf("read conf error %s", err.Error())
	}

	// 获取字符串类型的配置项的值
	fmt.Printf("String: %+v\n", viper.GetString("StringKey"))
	// 获取整数类型的配置项的值
	fmt.Printf("Int: %+v\n", viper.GetInt("IntKey"))
	// 获取浮点数类型的配置项的值
	fmt.Printf("Float64: %+v\n", viper.GetFloat64("Float64Key"))
	// 获取布尔类型的配置项的值
	fmt.Printf("Bool: %+v\n", viper.GetBool("BoolKey"))
	// 获取Int切片类型的配置项的值
	fmt.Printf("IntSlice: %+v\n", viper.GetIntSlice("IntSliceKey"))
	// 获取Interface类型的配置项的值
	fmt.Printf("Map: %+v\n", viper.Get("MapKey"))
	// 获取映射类型的配置项的值
	fmt.Printf("Map: %+v\n", viper.GetStringMap("MapKey"))

	// 获取映射类型的配置项的值到结构体
	type MapKey struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
	var mk MapKey
	viper.UnmarshalKey("MapKey", &mk)
	fmt.Printf("Map: %+v\n", mk)

	// 获取所有配置项
	settings := viper.AllSettings()
	for key, value := range settings {
		fmt.Printf("%s=%v\n", key, value)
	}

	// 绑定环境变量 GOPATH
	viper.BindEnv("GOPATH")
	fmt.Printf("GOPATH: %+v\n", viper.Get("GOPATH"))

	// 绑定环境变量 GOROOT 到 root 这个Key上
	viper.BindEnv("root", "GOROOT")
	fmt.Printf("GOROOT: %+v\n", viper.Get("root"))
}
