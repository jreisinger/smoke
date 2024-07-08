package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type filesPresent []string

func FilesPresent(hostName string, config []byte) (string, bool) {
	var files filesPresent
	if err := json.Unmarshal(config, &files); err != nil {
		return fmt.Sprintf("unmarshal FilesPresent config: %v", err), false
	}

	cmd := fmt.Sprintf("ls -l %s", strings.Join(files, " "))
	_, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), false
	}

	return strings.Join(files, ", "), true
}
