// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func useGzip(engine *gin.Engine) {
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	engine := gin.New()
	// engine.Use(gzip.Gzip(gzip.DefaultCompression))  //如果需要开启gzip压缩，取消这一行的注释
	engine.Handle("POST", "/query", downloadFile)
	engine.Handle("GET", "/", homepage)
	engine.Run(":8080")
}

func downloadFile(ctx *gin.Context) {
	// 此处省略查询的业务逻辑
	//  todo:
	// 下面开始下载的准备
	ctx.Writer.WriteHeader(200)
	ctx.Header("Content-Type", "text/plain; charset=utf-8")
	ctx.Header("Transfer-Encoding", "chunked") // 告诉浏览器，分段的流式的输出数据
	//   ctx.Header("Content-Encoding", "gzip") // 输出不是gzip内容，又加上这个头，浏览器会拒收。这里是个实验，不要加这行代码
	now := time.Now()
	fileName := now.Format("20060102_150405.csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName)) // 设置下载的文件名
	ctx.Writer.WriteHeaderNow()
	// 下面模拟一个周期非常长的数据处理和下载过程
	for i := 0; i < 100; i++ {
		ctx.Writer.WriteString("\"")
		ctx.Writer.WriteString("hahah")
		ctx.Writer.WriteString("\"\t")
		ctx.Writer.WriteString("\"")
		ctx.Writer.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		ctx.Writer.WriteString("\"\n")
		ctx.Writer.Flush() // 产生一定的数据后， flush到浏览器端
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

func homepage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.Writer.WriteString(`
<html>
<body>
open window and to download:
<a href="javascript:download()">download</a>
<script>
function download(){
    var handle = window.open("about:blank", "my_download_window");
	document.forms[0].target = "my_download_window";
	document.forms[0].json.value="ahfu test";
	document.forms[0].submit();
}
</script>
<form action="/query" method="POST" enctype="multipart/form-data">
<input type="hidden" name="json" value=""/>
</form>
</body>
</html>
`)
}
