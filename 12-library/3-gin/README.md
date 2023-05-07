# Gin基础

## 入门Demo

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Golang Tutorial!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// curl --location --request GET '127.0.0.1:8080/hello'
```

## 请求方法

```go
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
```

## 响应格式

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

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
```

## 参数绑定

### 路径参数

```go
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

	// 确切的路由在参数路由之前解析，无论它们的定义顺序如何。
	// 以 /user/groups 开头的路由永远不会被匹配为 /user/:name/ 的路由
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/user/golang'
// curl --location --request GET '127.0.0.1:8080/user/golang/bac'
// curl --location --request GET '127.0.0.1:8080/user/groups'
// curl --location --request GET '127.0.0.1:8080/user/groups/bac'
```

### Get+Post 参数

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMethod(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func PostMethod(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

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

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
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
```

### 参数绑定

```go
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

	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	// Example for binding JSON ({"user": "manu", "password": "123"})
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

	// Example for binding a HTML form (user=manu&password=123)
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

	// function only binds the query params and not the post data. (user=manu&password=123)
	router.POST("/loginQuery", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
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
```

## 参数校验

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	// 包含绑定和验证的数据,bookabledate就是自定义的验证函数
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	// gtfield=CheckIn只对数字或时间有效，参考官网链接
	// https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// 自定义验证器
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) { // 输入的日期必须大于今天的日期，否则验证失败
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	// 注册验证器
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate.RegisterValidation("bookabledate", bookableDate)
	}
	// http://127.0.0.1:8085/bookable?check_in=2022-01-07&check_out=2022-01-08
	// https://frankhitman.github.io/zh-CN/gin-validator/
	router.GET("/bookable", getBookable)
	router.Run()
}

func getBookable(context *gin.Context) {
	var book Booking
	if err := context.ShouldBindWith(&book, binding.Query); err == nil {
		context.JSON(200, gin.H{"message": "book date is valid"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

// check_in=2022-01-11&check_out=2022-01-12 (输入的日期必须大于今天的日期，否则验证失败)
// curl --location --request GET 'http://127.0.0.1:8080/bookable?check_in=2022-01-11&check_out=2022-01-12'

```

## 文件上传

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("Filename")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./"+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

/*
curl --location --request POST '127.0.0.1:8080/upload' \
--form 'Filename=@"/Users/Pictures/psc.jpg"'
*/
```

## 路由分组

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
```

## 中间件

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 默认情况下创建没有任何中间件的路由器
	r := gin.New()

	// 全局中间件
	r.Use(gin.Logger())

	// 恢复中间件从任何恐慌中恢复，如果有，则写入 500。
	r.Use(gin.Recovery())

	// 添加中间件在context中添加一个key-value
	r.Use(func(c *gin.Context) {
		fmt.Println("I am middleware")
		c.Set("flag", "I am middleware")
	})

	r.Use(func(c *gin.Context) {
		fmt.Println("begin")
		c.Next()
		fmt.Println("end")
	})

	// Per route middleware, you can add as many as you desire.
	r.GET("/middleware", func(c *gin.Context) {
		c.String(http.StatusOK, c.Keys["flag"].(string))
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/middleware'
```

## 自定义日志

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	r := gin.New()

	logger, _ := zap.NewProduction()

	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			// log request ID
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
				fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))

			return fields
		}),
	}))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.POST("/ping", func(c *gin.Context) {
		c.Writer.Header().Add("X-Request-Id", "9012-5678-1234")
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// curl --location --request GET '127.0.0.1:8080/ping'
// curl --location --request POST '127.0.0.1:8080/ping'
```

## 优雅关闭

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 在 goroutine 中初始化服务器，以便它不会阻止下面的正常关闭处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号正常关闭服务器，超时为 5 秒。
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 上下文用于通知服务器它有 5 秒的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
```



## 参考

https://www.liwenzhou.com/posts/Go/zap-in-gin/