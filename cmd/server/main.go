package main

import (
	"dns-go/packet"
	"fmt"
	"net"
)

func main() {
	server, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 53})
	trace(err)
	fmt.Println("Local Addr:", server.LocalAddr().String())

	buff := make([]byte, 1024)
	for {
		n, remote, err := server.ReadFromUDP(buff)
		trace(err)

		fmt.Printf("%s -> %s\n", remote.String(), buff[:n])

		req := &packet.Request{}
		_, err = packet.UnmarshalRequest(buff[:n], req)
		trace(err)
		fmt.Printf("%+v\n", req)

		response, err := packet.NewResponse(req)
		trace(err)

		_, err = server.WriteToUDP(response, remote)
		trace(err)
	}
}

func trace(err error) {
	if err != nil {
		panic(err)
	}
}
