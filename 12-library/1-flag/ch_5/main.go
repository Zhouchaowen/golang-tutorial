package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

func main() {
	viper.SetConfigFile("12-library/1-flag/ch_5/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件失败的情况
		log.Panicf("read conf error %s", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		fmt.Printf("String: %+v\n", viper.GetString("StringKey"))
	})

	// 获取字符串类型的配置项的值
	fmt.Printf("String: %+v\n", viper.GetString("StringKey"))

	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt)

	<-stopper
}
