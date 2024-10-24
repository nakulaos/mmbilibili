package tools

import (
	"fmt"
	"net"
)

func SplitIPAndPort(address string) (string, string, error) {
	ip, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", "", fmt.Errorf("error splitting IP and port: %v", err)
	}
	return ip, port, nil
}
