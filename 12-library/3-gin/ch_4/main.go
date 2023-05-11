package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 此处理程序将匹配 /user/john 但不匹配 /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 这个会匹配 /user/john/ 还有 /user/john/send
	// 如果没有其他路由器匹配 /user/john, 它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 确切的路由在参数路由之前解析，无论它们的定义顺序如何，以 /user/groups 开头的路由永远不会被匹配为 /user/:name/ 的路由
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/user/golang'
// curl --location --request GET '127.0.0.1:8080/user/golang/bac'
// curl --location --request GET '127.0.0.1:8080/user/groups'
// curl --location --request GET '127.0.0.1:8080/user/groups/bac'
