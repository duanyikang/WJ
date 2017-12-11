package models

import (
	"net"
	"fmt"
	"time"
	"bufio"
	"WJ/models/codec"
	"WJ/src/github.com/astaxie/beego"
)

var quitSemaphore chan bool
var temp = "231231231231"
var summsg = ""

func ClientStart() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", beego.AppConfig.String("longdir"))
	conn, _ := net.DialTCP("tcp", nil, tcpAddr) //开启连接
	if conn != nil {
		defer conn.Close()
	}else {
		return
	}
	fmt.Println("Connected!")
	go onMessageRecived(conn) //接收消息
	go sendMessage(conn)
	<-quitSemaphore
}

func ClientSendmsg(data string) {
	temp = data;
}

func GetMessge()( string)  {
	return summsg
}

// 发送消息
func sendMessage(conn *net.TCPConn) {
	//发送消息
	for {
		time.Sleep(1 * time.Second)
		if temp == "quit" {
			quitSemaphore <- true
			break
		}
		//lk
		//b :=[]byte(msg +"\n")
		//处理加密
		b, _ := codec.Encode(temp + "\n")
		conn.Write(b)

	}
}


func onMessageRecived(conn *net.TCPConn){
	reader := bufio.NewReader(conn)
	for {
		//解密
		msg, err := codec.Decode(reader) //reader.ReadString('\n')
		fmt.Println("客户端收到的消息：", msg)
		temp= "231231231231"
		if err != nil {
			quitSemaphore <- true
			break
		}
		summsg+=msg+"------我是分割线-------";
	}
}
