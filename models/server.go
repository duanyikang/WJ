package models

import (
	"bufio"
	"fmt"
	"net"
	"WJ/models/codec"
	"strings"
	"WJ/src/github.com/astaxie/beego"
)

// 用来记录所有的客户端连接
var ConnMap map[string]*net.TCPConn

func StartLonglink()  {
	var tcpAddr *net.TCPAddr
	ConnMap = make(map[string]*net.TCPConn) //初始化
	tcpAddr,_=net.ResolveTCPAddr("tcp",beego.AppConfig.String("longdir"))

	tcpListener,_:=net.ListenTCP("tcp",tcpAddr) //开启tcp 服务
	//退出时关闭
	defer tcpListener.Close()
	fmt.Println("Long server start  ")
	for{
		tcpConn,err :=tcpListener.AcceptTCP()
		if err !=nil {
			continue
		}
		fmt.Println("A client connected : "+ tcpConn.RemoteAddr().String())
		// 新连接加入 map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn

		go tcpPipe(tcpConn)
	}

}

//处理发送过来的消息
func tcpPipe(conn *net.TCPConn)  {
	ipStr :=conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected : "+ ipStr)
		conn.Close()
	}()
	//读取数据
	reader :=bufio.NewReader(conn)
	for {
		message ,err :=codec.Decode(reader)//reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if(strings.Contains(message,"231231231231")){
			continue
		}
		fmt.Println("服务端收到的消息：",message)
		//这里返回消息改为广播
		boradcastMessage(conn.RemoteAddr().String()+":"+string(message))
	}
}
//广播给其它
func boradcastMessage(message string)  {
	//遍历所有客户端并发消息
	for _,conn :=range ConnMap{
		b,err :=codec.Encode(message)
		if err != nil {
			continue
		}
		conn.Write(b)
	}
}