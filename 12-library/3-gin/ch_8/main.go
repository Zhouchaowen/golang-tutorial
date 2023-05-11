package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 为多部分表单设置内存下限（默认值为 32 MiB）
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, _ := c.FormFile("Filename")
		log.Println(file.Filename)

		// 将文件上传到特定的 dst。
		c.SaveUploadedFile(file, "./"+file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

/*
	curl --location --request POST '127.0.0.1:8080/upload' \
	--form 'Filename=@"/Users/Pictures/psc.jpg"'
*/
