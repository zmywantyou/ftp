package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听
	// net.Listen函数创建一个TCP监听器，监听`localhost:8080`，等待连接
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()
	for {
		// 无限循环来接受客户端连接
		conn, _ := listener.Accept()
		// 对于每个连接开启一个go协程，处理连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	//获取文件大小
	// sizebyte := make([]byte, 4)
	// n2, _ := conn.Read(sizebyte)
	// sizeint := int32(binary.LittleEndian.Uint32(sizebyte[:n2]))
	// 读取文件名
	filname := make([]byte, 1024)
	n, _ := conn.Read(filname)
	// 处理数据并回复客户端
	if n != 0 {
		fmt.Printf("Received:'\t' %s", string(filname[:n]))
	}
	Filname := string(filname[:n])
	Filname = "D:\\Download\\桌面\\severs\\data\\" + Filname
	file, err := os.Create(Filname)

	fmt.Print("文件路径", Filname)
	if err != nil {
		fmt.Println("Failed to read from connection:", err)
		return
	}

	//读取文件内容
	systemReader(file, conn)

	//conn.Write([]byte("Message received."))
}
func systemReader(file *os.File, conn net.Conn) {
	defer conn.Close()
	defer file.Close()
	offer := int64(0)
	for {

		buffer := make([]byte, 1024)
		n, _err := conn.Read(buffer)
		if _err != nil {
			fmt.Print("文件接收完成")
			return
		}

		n2, err := file.WriteAt(buffer[:n], offer)
		offer = int64(n) + offer
		fmt.Println("文件段接受", n, "个字节", err)
		fmt.Println("文件段写入", n2, "个字节", err)
	}

}
