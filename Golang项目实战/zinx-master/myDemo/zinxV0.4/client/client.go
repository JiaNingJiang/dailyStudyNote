package main

import (
	"GO_Demo/zinx/znet"
	"fmt"
	"io"
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

		dp := znet.NewDataPack() //封包、解包对象
		//2.向服务器发送数据(按照TLV格式发送，即需要封包)
		sendMsg := znet.NewMessage(0, []byte("Hello ZINX")) //构造Message格式消息
		binarySendMsg, _ := dp.Pack(sendMsg)                //进行封包
		if _, err := conn.Write(binarySendMsg); err != nil {
			fmt.Println("[client] send message to server is err:", err)
			break
		}
		//3.等待服务器的回复(TLV格式的二进制Message包)
		//3.1 读取Message消息头
		binaryData := make([]byte, dp.MsgHeadLen())
		if _, err := io.ReadFull(conn, binaryData); err != nil {
			fmt.Println("[client] recv head data is err:", err)
			break
		}
		fmt.Println("binaryData:", binaryData)
		//3.2 对消息头进行解包
		msgData, err := dp.UnPack(binaryData)
		if err != nil {
			fmt.Println("[client] Unpack binary data is err", err)
			break
		}
		//3.3 按照消息头中消息长度读取消息实体部分
		bodyData := make([]byte, msgData.GetMsgLenth())
		if _, err := io.ReadFull(conn, bodyData); err != nil { //读取消息实体
			fmt.Println("[client] recv body data is err:", err)
			break
		}
		msgData.SetMsgData(bodyData)
		fmt.Printf("[client] recv data from server , msgID:%v , msg:%s\n", msgData.GetMsgID(), string(msgData.GetMsgData()))

		time.Sleep(time.Second)
	}

}
