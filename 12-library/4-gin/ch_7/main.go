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

func getBookable(context *gin.Context) {
	var book Booking
	if err := context.ShouldBindQuery(&book); err == nil {
		context.JSON(200, gin.H{"message": "book date is valid"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

type Login struct {
	UserName string `form:"user_name" binding:"required"`
	PassWord string `form:"pass_word" binding:"required,min=8"` // 密码必须大于8位
}

func LoginHandler(context *gin.Context) {
	var login Login
	if err := context.ShouldBindQuery(&login); err == nil {
		context.JSON(200, gin.H{"message": "lgoin date is valid"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

func main() {
	router := gin.Default()

	router.GET("/login", LoginHandler)

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

// curl --location --request GET 'http://127.0.0.1:8080/login?user_name=zcw&pass_word=asdasda'
// curl --location --request GET 'http://127.0.0.1:8080/login?user_name=zcw&pass_word=asdasdas'

// check_in=2022-01-11&check_out=2022-01-12 (输入的日期必须大于今天的日期，否则验证失败)
// curl --location --request GET 'http://127.0.0.1:8080/bookable?check_in=2022-01-11&check_out=2022-01-12'
