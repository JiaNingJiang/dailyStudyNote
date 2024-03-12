package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	//1.连接到服务器
	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("client start is err:%s , exit!\n", err)
	}

	for {
		//2.向服务器发送数据
		if _, err := conn.Write([]byte("Hello ZINX")); err != nil {
			fmt.Println("write conn is err :", err)
			continue
		}
		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("recv conn is err :", err)
			continue
		}
		fmt.Printf("server call back is :%s\n", string(buf[:count]))

		time.Sleep(time.Second)
	}

}
