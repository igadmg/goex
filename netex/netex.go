package netex

import (
	"fmt"
	"net"
)

func GetOutboundIP() (net.IP, error) {
	// 8.8.8.8:80 doesn't need to be reachable or even exist.
	// This just helps Go determine the local IP for external traffic.
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return net.IP{}, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IP{}, err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return net.IP{}, fmt.Errorf("no external interface found")
}
