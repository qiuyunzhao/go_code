package main

import (
	"bufio"
	"fmt"
	"os"
)

//创建一个新文件，写入内容 5句 "hello, Gardon"
func main() {

	//filePath := "./18_文件操作/sources/abc.txt" // “./” 表示当前工程路径，可不加

	filePath := "./18_文件操作/sources1/abc.txt" // 文件夹不存在会报错The system cannot find the path specified. 可以递归创建文件夹来解决
	// 递归创建文件夹
	err := os.MkdirAll("./18_文件操作/sources1", os.ModePerm)
	if err != nil {
		fmt.Printf("os.MkdirAll err=%v\n", err)
		return
	}

	//OpenFile(name string, flag int, perm FileMode)   flag：操作模式    perm：操作权限linux有用
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//写入5句 "hello, Gardon"
	str := "hello,Gardon\r\n" // \r\n 表示换行

	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {

		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实内容是先写入到缓存的,
	// 所以需要调用Flush方法，将缓冲的数据真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}
