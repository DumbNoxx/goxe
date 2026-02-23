package filters

import (
	"bytes"
	"net"
)

func GetIpBurstDetection(log []byte) string {
	tokens := bytes.Fields(log)

	var foundIp net.IP

	for _, token := range tokens {
		if ip := net.ParseIP(string(token)); ip != nil {
			foundIp = ip
			break
		}
		if ip, _, err := net.ParseCIDR(string(token)); err == nil {
			foundIp = net.IP(ip)
			break
		}
	}

	return foundIp.String()
}
