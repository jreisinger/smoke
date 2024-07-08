package tests

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)

type portsOpen []string

func PortsOpen(hostName string, config []byte) (string, error) {
	var op portsOpen
	if err := json.Unmarshal(config, &op); err != nil {
		return "", fmt.Errorf("unmarshal OpenPorts config: %v", err)
	}

	for _, port := range op {
		addr := net.JoinHostPort(hostName, port)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			return fmt.Sprintf("can't connect to %s", addr), err
		}
		conn.Close()
	}
	return strings.Join(op, ", "), nil
}
