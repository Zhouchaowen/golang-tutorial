package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 默认情况下创建没有任何中间件的路由器
	r := gin.New()

	// 全局日志中间件
	r.Use(gin.Logger())

	// 全局恢复中间件从任何恐慌中恢复，如果有，则写入 500。
	r.Use(gin.Recovery())

	// 自定义中间件，在context中添加一个key-value
	r.Use(func(c *gin.Context) {
		fmt.Println("I am middleware")
		c.Set("flag", "I am middleware")
	})

	r.Use(func(c *gin.Context) {
		fmt.Println("begin")
		c.Next()
		fmt.Println("end")
	})

	// 每个路由中间件，您可以根据需要添加任意数量的中间件。
	r.GET("/middleware", func(c *gin.Context) {
		c.String(http.StatusOK, c.Keys["flag"].(string)) // 从上下文读取flag的value
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/middleware'
