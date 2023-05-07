package main

import "github.com/gin-gonic/gin"

// https://www.runoob.com/http/http-methods.html
func main() {
	// 使用默认中间件创建一个gin路由器
	router := gin.Default()

	router.GET("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	router.POST("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	router.PUT("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	router.DELETE("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})
	router.PATCH("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PATCH",
		})
	})
	router.HEAD("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "HEAD",
		})
	})
	router.OPTIONS("/some", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "OPTIONS",
		})
	})

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port
}

// curl --location --request GET '127.0.0.1:8080/some'
// curl --location --request POST '127.0.0.1:8080/some'
// curl --location --request PUT '127.0.0.1:8080/some'
// curl --location --request DELETE '127.0.0.1:8080/some'
// curl --location --request PATCH '127.0.0.1:8080/some'
// curl --location --request HEAD '127.0.0.1:8080/some'
// curl --location --request OPTIONS '127.0.0.1:8080/some'
