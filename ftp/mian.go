package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// net.Dial函数连接 TCP 服务端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	fmt.Print("已经连接")
	// 延迟关闭连接
	defer conn.Close()
	// 向服务器发送文件名
	var filsrc string
	filname := make([]byte, 1024)
	fmt.Scan(&filsrc)
	fmt.Scan(&filname)
	conn.Write(filname)
	handFile(filsrc, conn)

}

func handFile(Filsrc string, conn net.Conn) {
	var size int
	file, err := os.Open(Filsrc)
	defer file.Close()
	defer conn.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	for {

		buff := make([]byte, 1024)
		n, err := file.Read(buff)
		size = n + size
		if err == io.EOF {

			//fmt.Print("传输中断")
			break
		}
		conn.Write(buff[:n])
		fmt.Println("文件段传输成功，大小为", n)
	}

	fmt.Println("文件传输成功")
	fmt.Println("文件大小为：", size, "字节")
}
