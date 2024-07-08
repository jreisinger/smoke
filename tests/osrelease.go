package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type osRelease struct {
	ID         string
	VERSION_ID string
}

func OsRelease(hostName string, config []byte) (string, bool) {
	var os osRelease
	if err := json.Unmarshal(config, &os); err != nil {
		return fmt.Sprintf("unmarshal OsRelease config: %v", err), false
	}

	cmd := "cat /etc/os-release"
	out, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), false
	}
	for _, line := range strings.Split(string(out), "\n") {
		parts := strings.Split(line, "=")
		if parts[0] == "ID" && parts[1] != os.ID {
			msg := fmt.Sprintf("want ID=%s, got ID=%s", os.ID, parts[1])
			return msg, false
		}
		if parts[0] == "VERSION_ID" && parts[1] != os.VERSION_ID {
			msg := fmt.Sprintf("want VERSION_ID=%s, got VERSION_ID=%s", os.VERSION_ID, parts[1])
			return msg, false
		}
	}

	return fmt.Sprintf("ID=%s, VERSION_ID=%s", os.ID, os.VERSION_ID), true
}
