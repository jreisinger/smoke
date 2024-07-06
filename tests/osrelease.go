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

func OsRelease(host string, config []byte) (string, error) {
	var os osRelease
	err := json.Unmarshal(config, &os)
	if err != nil {
		return "", fmt.Errorf("unmarshal OsRelease config: %v", err)
	}

	out, err := helper.Ssh(host, "cat /etc/os-release")
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(out), "\n") {
		parts := strings.Split(line, "=")
		if parts[0] == "ID" && parts[1] != os.ID {
			out := fmt.Sprintf("want ID=%s, got ID=%s", os.ID, parts[1])
			return out, nil
		}
		if parts[0] == "VERSION_ID" && parts[1] != os.VERSION_ID {
			out := fmt.Sprintf("want VERSION_ID=%s, got VERSION_ID=%s", os.VERSION_ID, parts[1])
			return out, nil
		}
	}
	return "", nil
}
