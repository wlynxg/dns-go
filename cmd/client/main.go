package main

import (
	"dns-go/packet"
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(114, 114, 114, 114),
		Port: 53,
	})
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	defer socket.Close()

	sendData := packet.NewRequest("google.com")
	fmt.Println("send data:", sendData)
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err: ", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err: ", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
