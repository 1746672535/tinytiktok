package tools

import (
	"net"
)

func GetFreePort() int {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")
	TCPListener, _ := net.ListenTCP("tcp", addr)
	defer TCPListener.Close()
	return TCPListener.Addr().(*net.TCPAddr).Port
}

func GetLocalIP() string {
	ipAddr, _ := net.ResolveIPAddr("ip", "www.baidu.com")
	rAddr, err1 := net.ResolveIPAddr("ip4:icmp", ipAddr.String())
	lAddr, err2 := net.ResolveIPAddr("ip4:icmp", "")
	con, err3 := net.DialIP("ip4:icmp", lAddr, rAddr)
	if err1 != nil || err2 != nil || err3 != nil {
		return ""
	}
	defer con.Close()
	return con.LocalAddr().String()
}
