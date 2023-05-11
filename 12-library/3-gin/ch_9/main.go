package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由分组

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 login")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 submit")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 read")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 login")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 submit")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 read")
		})
	}

	router.Run(":8080")
}

// curl --location --request POST '127.0.0.1:8080/v1/login'
// curl --location --request POST '127.0.0.1:8080/v2/login'
