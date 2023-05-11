package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMethod(c *gin.Context) {
	// 解析路由参数 firstname 没有该参数将所有默认值 Guest
	// 将解析URL /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func PostMethod(c *gin.Context) {
	// 解析POST 提交的数据
	nick := c.DefaultPostForm("nick", "anonymous")
	message := c.PostForm("message")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func GetPostMethod(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.JSON(200, gin.H{
		"id":      id,
		"page":    page,
		"message": message,
		"nick":    name,
	})
}

func main() {
	router := gin.Default()

	router.GET("/welcome", GetMethod)

	router.POST("/form_post", PostMethod)

	router.POST("/post", GetPostMethod)

	router.Run(":8080")
}

/*
curl --location --request GET '127.0.0.1:8080/welcome?firstname=Jane&lastname=Doe'

curl --location --request POST '127.0.0.1:8080/form_post' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'message=123' \
--data-urlencode 'nick=zcw'

curl --location --request POST '127.0.0.1:8080/post?id=1&page=10' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'message=123' \
--data-urlencode 'nick=zcw'
*/
