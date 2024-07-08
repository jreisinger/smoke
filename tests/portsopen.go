package tests

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)

type portsOpen []string

func PortsOpen(hostName string, config []byte) (string, bool) {
	var op portsOpen
	if err := json.Unmarshal(config, &op); err != nil {
		return fmt.Sprintf("unmarshal OpenPorts config: %v", err), false
	}

	for _, port := range op {
		addr := net.JoinHostPort(hostName, port)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			return fmt.Sprintf("can't connect to %s", addr), false
		}
		conn.Close()
	}
	return strings.Join(op, ", "), true
}
