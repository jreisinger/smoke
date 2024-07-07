package tests

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type openPorts []string

func OpenPorts(hostName string, config []byte) (string, error) {
	var op openPorts
	err := json.Unmarshal(config, &op)
	if err != nil {
		return "", fmt.Errorf("unmarshal OpenPorts config: %v", err)
	}

	for _, port := range op {
		addr := net.JoinHostPort(hostName, port)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			return fmt.Sprintf("can't connect to %s", addr), nil
		}
		conn.Close()
	}
	return "", nil
}
