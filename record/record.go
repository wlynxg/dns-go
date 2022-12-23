package record

import (
	"dns-go/packet"
	"errors"
	"net"
	"sync"
)

var (
	data = map[string][]net.IP{
		"baidu.com": {
			net.ParseIP("39.156.66.10"),
			net.ParseIP("110.242.68.66"),
		},
	}

	nameservers = []string{
		"114.114.114.114",
		"8.8.8.8",
		"223.5.5.5",
		"223.6.6.6",
	}

	conn *net.UDPConn
	once sync.Once
)

func QueryA(domain string, qtype packet.QueryType) ([]net.IP, error) {
	if ip, ok := data[domain]; ok {
		return ip, nil
	} else {
		return nil, errors.New("unknown domain")
	}
}

func QueryByOtherNameserver(domain string, qtype packet.QueryType) {
	once.Do(func() {
		var err error
		conn, err = net.ListenUDP("udp", nil)
		if err != nil {
			panic(err)
		}
	})

	var (
		results []net.IP
		buff    = make([]byte, 1024)
		data    = packet.NewRequest(domain, qtype)
	)

	for _, nameserver := range nameservers {
		_, err := conn.WriteToUDP(data, &net.UDPAddr{IP: net.ParseIP(nameserver), Port: 53})
		if err != nil {
			continue
		}

		n, _, err := conn.ReadFromUDP(buff)
		if err != nil {
			continue
		}

		res := new(packet.Response)
		_, err = packet.UnmarshalResponse(buff[:n], res)
		if err != nil {
			continue
		}

		results = append(results, res.Answers)
	}
	return
}
