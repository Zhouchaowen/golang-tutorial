package main

import "github.com/gin-gonic/gin"

// https://www.runoob.com/http/http-methods.html
func main() {
	// 使用默认中间件创建一个gin路由器
	router := gin.Default()

	// GET 方法请求
	router.GET("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	// POST 方法请求
	router.POST("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	// PUT 方法请求
	router.PUT("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	// DELETE 方法请求
	router.DELETE("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	// PATCH 方法请求
	router.PATCH("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PATCH",
		})
	})

	// HEAD 方法请求
	router.HEAD("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "HEAD",
		})
	})

	// OPTIONS 方法请求
	router.OPTIONS("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "OPTIONS",
		})
	})

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// 自己定义启动端口 3000
	// router.Run(":3000")
}

// curl --location --request GET '127.0.0.1:8080/some'
// curl --location --request POST '127.0.0.1:8080/some'
// curl --location --request PUT '127.0.0.1:8080/some'
// curl --location --request DELETE '127.0.0.1:8080/some'
// curl --location --request PATCH '127.0.0.1:8080/some'
// curl --location --request HEAD '127.0.0.1:8080/some'
// curl --location --request OPTIONS '127.0.0.1:8080/some'
