package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Person Bind uri
type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// Login Bind form，json，xml，header
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 将绑定URI参数示例 xx.xx.xx.xx:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	// 绑定 JSON 的示例 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定 HTML 表单的示例 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 函数仅绑定查询参数，而不绑定发布数据。 (user=manu&password=123)
	router.POST("/loginQuery", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		// c.ShouldBind(&x)
		if err := c.ShouldBindQuery(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

/*
curl --location --request GET '127.0.0.1:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3'

curl --location --request POST '127.0.0.1:8080/loginJSON' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user":"manu",
    "password":"123"
}'

curl --location --request POST '127.0.0.1:8080/loginForm' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'user=manu' \
--data-urlencode 'password=123'

curl --location --request POST '127.0.0.1:8080/loginQuery?user=manu&password=123'
*/
