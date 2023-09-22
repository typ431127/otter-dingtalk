package ext

import (
	"net"
	"os"
	"otter-dingtalk/internal/global"
)

const defaultip = "127.0.0.1"

func InterfaceAddrs() string {
	addr, err := net.InterfaceByName(global.GL_INTERFACE)
	if err != nil {
		global.GL_LOG.Warnln(err)
		return defaultInterface()
	}
	addrs, err := addr.Addrs()
	if err != nil {
		global.GL_LOG.Warnln(err)
		return defaultInterface()
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return defaultInterface()
}

func defaultInterface() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return defaultip
}

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		global.GL_LOG.Warnln(err)
		return defaultip
	}
	return hostname
}

func Ipv4Parse(ip string) bool {
	if net.ParseIP(ip) == nil {
		//global.GL_LOG.Errorf("ip:%s校验失败", ip)
		return false
	}
	return true
}
