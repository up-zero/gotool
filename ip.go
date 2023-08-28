package gotool

import "net"

// Ipv4sLocal 获取本地ipv4地址
func Ipv4sLocal() ([]string, error) {
	res := make([]string, 0)
	// 获取所有网卡地址
	address, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	// 遍历网卡地址
	for _, addr := range address {
		// 获取ip地址
		ip := addr.(*net.IPNet).IP
		// 过滤ip地址
		// 127.0.0.0 - 127.255.255.255
		// 169.254.0.0 - 169.254.255.255
		// 169.254.0.0 - 169.254.255.255
		if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
			continue
		}
		// 过滤ipv6地址
		if ip.To4() != nil {
			res = append(res, ip.String())
		}
	}

	return res, nil
}
