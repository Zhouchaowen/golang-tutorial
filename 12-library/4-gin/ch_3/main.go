package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据, gin.H 实践就是一个map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// 返回一个结构体，将会编码为Json格式
	r.GET("/moreJSON", func(c *gin.Context) {
		// 定义返回的 struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意其中 msg.Name 在 JSON 中变为“user”, 应为指定了Json tag
		// 返回结果: {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	// 返回XML格式数据
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// 返回YAML数据
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/someJSON'
// curl --location --request GET '127.0.0.1:8080/moreJSON'
// curl --location --request GET '127.0.0.1:8080/someXML'
// curl --location --request GET '127.0.0.1:8080/someYAML'
