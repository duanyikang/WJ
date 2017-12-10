package models

import (
	"net"
	"fmt"
	"time"
	"bufio"
	"WJ/models/codec"
)

var quitSemaphore chan bool
var conn *net.TCPConn

func ClientSendMsg() {

}

func ClientStart() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ = net.DialTCP("tcp", nil, tcpAddr) //开启连接
	defer conn.Close()                         //关闭连接
	fmt.Println("Connected!")
	go onMessageRecived(conn) //接收消息
	<-quitSemaphore
}

func ClientSendmsg(msg string) {
	go sendMessage(conn, msg)
}

// 发送消息
func sendMessage(conn *net.TCPConn, msg string) {
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
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		//解密
		msg, err := codec.Decode(reader) //reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
