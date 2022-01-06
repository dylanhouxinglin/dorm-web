package utils

import (
	"fmt"
	"net"
)

func AcquireLocalHost() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				host := ipnet.IP.String()
				fmt.Printf("成功获取本机host: %s\n", host)
				return host
			}
		}
	}
	return ""
}
