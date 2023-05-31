package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	c.Next()
}

func checkChunk(c *gin.Context) {
	hash := c.Query("hash")
	hashPath := fmt.Sprintf("./uploadFile/%s", hash)
	chunkList := []string{}
	isExistPath, err := PathExists(hashPath)
	if err != nil {
		fmt.Println("获取hash路径错误", err)
	}

	if isExistPath {
		files, err := ioutil.ReadDir(hashPath)
		state := 0
		if err != nil {
			fmt.Println("文件读取错误", err)
		}
		for _, f := range files {
			fileName := f.Name()
			chunkList = append(chunkList, fileName)
			fileBaseName := strings.Split(fileName, ".")[0]
			if fileBaseName == hash {
				state = 1
			}
		}

		c.JSON(200, gin.H{
			"state":     state,
			"chunkList": chunkList,
		})
	} else {
		c.JSON(200, gin.H{
			"state":     0,
			"chunkList": chunkList,
		})
	}
}

func uploadChunk(c *gin.Context) {
	fileHash := c.PostForm("hash")
	file, err := c.FormFile("file")
	hashPath := fmt.Sprintf("./uploadFile/%s", fileHash)
	if err != nil {
		fmt.Println("获取上传文件失败", err)
	}

	isExistPath, err := PathExists(hashPath)
	if err != nil {
		fmt.Println("获取hash路径错误", err)
	}

	if !isExistPath {
		os.Mkdir(hashPath, os.ModePerm)
	}

	err = c.SaveUploadedFile(file, fmt.Sprintf("./uploadFile/%s/%s", fileHash, file.Filename))
	if err != nil {
		c.String(400, "0")
		fmt.Println(err)
	} else {
		chunkList := []string{}
		files, err := ioutil.ReadDir(hashPath)
		if err != nil {
			fmt.Println("文件读取错误", err)
		}
		for _, f := range files {
			fileName := f.Name()

			if f.Name() == ".DS_Store" {
				continue
			}
			chunkList = append(chunkList, fileName)
		}

		c.JSON(200, gin.H{
			"chunkList": chunkList,
		})
	}
}

func mergeChunk(c *gin.Context) {
	hash := c.Query("hash")
	fileName := c.Query("fileName")
	hashPath := fmt.Sprintf("./uploadFile/%s", hash)

	isExistPath, err := PathExists(hashPath)
	if err != nil {
		fmt.Println("获取hash路径错误", err)
	}

	if !isExistPath {
		c.JSON(400, gin.H{
			"message": "文件夹不存在",
		})
		return
	}
	isExistFile, err := PathExists(hashPath + "/" + fileName)
	if err != nil {
		fmt.Println("获取hash路径文件错误", err)
	}
	fmt.Println("文件是否存在", isExistFile)
	if isExistFile {
		c.JSON(200, gin.H{
			"fileUrl": fmt.Sprintf("http://127.0.0.1:9999/uploadFile/%s/%s", hash, fileName),
		})
		return
	}

	files, err := ioutil.ReadDir(hashPath)
	if err != nil {
		fmt.Println("合并文件读取失败", err)
	}
	complateFile, err := os.Create(hashPath + "/" + fileName)
	defer complateFile.Close()
	for _, f := range files {
		//.DS_Store
		//file, err := os.Open(hashPath + "/" + f.Name())
		//if err != nil {
		//	fmt.Println("文件打开错误", err)
		//}

		if f.Name() == ".DS_Store" {
			continue
		}

		fileBuffer, err := ioutil.ReadFile(hashPath + "/" + f.Name())
		if err != nil {
			fmt.Println("文件打开错误", err)
		}
		complateFile.Write(fileBuffer)
	}

	c.JSON(200, gin.H{
		"fileUrl": fmt.Sprintf("http://127.0.0.1:9999/uploadFile/%s/%s", hash, fileName),
	})

}

func main() {
	router := gin.Default()

	router.Use(Cors)
	router.GET("/checkChunk", checkChunk)

	router.POST("/uploadChunk", uploadChunk)

	router.GET("/mergeChunk", mergeChunk)

	router.Run("127.0.0.1:9999")
}

// https://juejin.cn/post/6844904159372656654
