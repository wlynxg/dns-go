package main

import (
	"dns-go/packet"
	"fmt"
	"log"
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
	n, _, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err: ", err)
		return
	}

	fmt.Printf("%v\n", data[:n])
	var offset int
	header := &packet.Header{}
	hs, err := packet.UnmarshalHeader(data[offset:n], header)
	if err != nil {
		log.Println(err)
		return
	}
	offset += hs
	fmt.Printf("%+v\n", header)

	queries := &packet.Queries{}
	qs, err := packet.UnmarshalQueries(data[offset:n], queries)
	if err != nil {
		log.Println(err)
		return
	}
	offset += qs
	fmt.Printf("%+v\n", queries)

	answers := &packet.Answers{}
	_, err = packet.UnmarshalAnswers(data[offset:n], answers)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", answers)
}
