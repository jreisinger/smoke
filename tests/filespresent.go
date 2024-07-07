package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type filesPresent []string

func FilesPresent(hostName string, config []byte) (string, error) {
	var files filesPresent
	if err := json.Unmarshal(config, &files); err != nil {
		return "", fmt.Errorf("unmarshal FilesPresent config: %v", err)
	}

	cmd := fmt.Sprintf("ls -l %s", strings.Join(files, " "))
	_, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), nil
	}

	return "", nil
}
