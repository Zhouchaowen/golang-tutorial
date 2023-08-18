package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Create 创建一个文件
func Create() {
	newFile, err := os.Create("golang-tutorial.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	fmt.Println(newFile)

	n, err := newFile.WriteString("golang tutorial.")
	if err != nil {
		panic(err)
	}
	fmt.Println("write", n, "byte to file")
}

// BufferWrite 缓存写
func BufferWrite() {
	// 打开文件，只写
	file, err := os.OpenFile("golang-tutorial1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{'g', 'o'},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)

	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unFlushedBufferSize := bufferedWriter.Buffered()
	fmt.Printf("Bytes buffered: %d\n", unFlushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)

	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)
}

// Read 读取文件内容
func Read() {
	// 打开文件，只读
	file, err := os.Open("golang-tutorial.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了,读取到文件会返回io.EOF的error。
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of bytes read: %d\n", bytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)

	// file.Read()可以读取一个下与len(b)字节的文件
	// io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSliceFull := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSliceFull)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSliceFull)
}

func ReadBuffer() {
	// 打开文件，创建buffered reader
	file, err := os.Open("golang-tutorial.txt")
	if err != nil {
		panic(err)
	}
	bufferedReader := bufio.NewReader(file)

	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read string: %s\n", dataString)

	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}

// FileInfo 获取文件信息
func FileInfo() {
	// 如果文件不存在，则返回错误
	fileInfo, err := os.Stat("golang-tutorial.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System info: %+v\n", fileInfo.Sys())
}

func main() {
	BufferWrite()
}
