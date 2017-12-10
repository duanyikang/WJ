package models

import (
	"net"
	"fmt"
	"time"
	"bufio"
	"WJ/models/codec"
)

var quitSemaphore chan bool
var msg="231231231231"

func ClientStart() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "47.95.200.181:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr) //开启连接
	defer conn.Close()                         //关闭连接
	fmt.Println("Connected!")
	go onMessageRecived(conn) //接收消息
	go sendMessage(conn)
	<-quitSemaphore
}

func ClientSendmsg(data string) {
	msg=data;
	fmt.Println(msg)
}

// 发送消息
func sendMessage(conn *net.TCPConn) {
	//发送消息
	for {
		time.Sleep(1 * time.Second)
		if msg == "quit" {
			quitSemaphore <- true
			break
		}
		//lk
		//b :=[]byte(msg +"\n")
		//处理加密
		b, _ := codec.Encode(msg + "\n")
		conn.Write(b)
		msg="231231231231"
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		//解密
		msg, err := codec.Decode(reader) //reader.ReadString('\n')
		fmt.Println("客户端收到的消息：",msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
