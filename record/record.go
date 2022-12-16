package record

import (
	"errors"
	"net"
)

var (
	data = map[string]net.IP{
		"baidu.com": net.IPv4(110, 242, 68, 66),
	}
)

func Query(domain string) (net.IP, error) {
	if ip, ok := data[domain]; ok {
		return ip, nil
	} else {
		return nil, errors.New("unknown domain")
	}
}
