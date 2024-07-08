package tests

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type resolvesToAddrs []string

func ResolvesToAddrs(hostName string, config []byte) (string, bool) {
	var wantAddrs resolvesToAddrs
	if err := json.Unmarshal(config, &wantAddrs); err != nil {
		return fmt.Sprintf("unmarshal ResolvesTo config: %v", err), false
	}

	addrs, err := net.LookupHost(hostName)
	if err != nil {
		return fmt.Sprintf("lookup %s: %v", hostName, err), false
	}
	if !helper.StringSlicesEqual(wantAddrs, addrs) {
		return fmt.Sprintf("want %s, got %s",
			strings.Join(wantAddrs, " "), strings.Join(addrs, " ")), false
	}

	return strings.Join(addrs, ", "), true
}
