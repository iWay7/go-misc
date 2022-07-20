package misc

import (
	"bytes"
	"net"
)

func CheckIPAddressInRange(start, end, target string) bool {
	startIPAddr := net.ParseIP(start)
	endIPAddr := net.ParseIP(end)
	targetIPAddr := net.ParseIP(target)
	return bytes.Compare(startIPAddr, targetIPAddr) <= 0 && bytes.Compare(endIPAddr, targetIPAddr) >= 0
}
