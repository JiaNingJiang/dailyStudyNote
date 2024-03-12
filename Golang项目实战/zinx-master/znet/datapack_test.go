package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	//模拟服务器
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("[server] listen is err:", err)
		return
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[server] accept is error:", err)
			return
		}
		go func(conn net.Conn) {
			//定义一个拆包的对象(拆包工具)
			dp := NewDataPack()
			for {
				//1.第一次读取，只读取消息头
				headData := make([]byte, dp.MsgHeadLen()) //存储缓冲区
				_, err := io.ReadFull(conn, headData)     //从连接套接字中读取指定长度的比特数存储到缓冲区
				if err != nil {
					fmt.Println("[server] read msg head is err:", err)
					break
				}
				//2.进行解包
				msgHead, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("[server] unpack data is err:", err)
					return
				}
				//3.第二次读取，根据消息头中的消息长度进行读取
				if msgHead.GetMsgLenth() > 0 { //说明当前接收的客户端Message的数据不为空，需要进行二次读取
					msg, _ := msgHead.(*Message)               //类型断言
					msg.Data = make([]byte, msg.GetMsgLenth()) //Data字段需要自行申请空间，大小等于消息长度指定的大小
					_, err := io.ReadFull(conn, msg.Data)      //此时的msg.Data是二进制小端字节序格式
					if err != nil {
						fmt.Println("server read msg is err:", err)
						return
					}
					fmt.Printf("[server]recv data:%s , dataID:%d , dataLen:%d\n", string(msg.Data), msg.ID, msg.Len)
				}

			}
		}(conn)

	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("[client] client connection is failed!")
	}
	//定义一个封包对象(封包工具)
	dp := NewDataPack()
	//封装第一个msg1包
	msg1 := &Message{
		ID:   1,
		Len:  4,
		Data: []byte{'a', 'b', 'c', 'd'},
	}
	packet1, err := dp.Pack(msg1) //对msg1进行封包
	if err != nil {
		fmt.Println("[client] Message1 Pack is err:", err)
		return
	}
	//封装第二个msg2包
	msg2 := &Message{
		ID:   2,
		Len:  5,
		Data: []byte{'h', 'e', 'l', 'l', 'o'},
	}
	packet2, err := dp.Pack(msg2) //对msg1进行封包
	if err != nil {
		fmt.Println("[client] Message2 Pack is err:", err)
		return
	}
	//将两个包粘在一起进行发送
	packet1 = append(packet1, packet2...)
	conn.Write(packet1)

	select {} //阻塞
}
