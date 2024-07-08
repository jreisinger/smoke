package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type helmReleases int

func HelmReleases(hostName string, config []byte) (string, error) {
	var hr helmReleases
	if err := json.Unmarshal(config, &hr); err != nil {
		return "", fmt.Errorf("unmarshal HelmReleases config: %v", err)
	}

	cmd := "helm ls -A"
	out, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), err
	}

	lines := strings.Split(string(out), "\n")
	releases := len(lines) - 2 // don't count header line and last empty line
	if releases != int(hr) {
		out := fmt.Sprintf("want count %d, got count %d", hr, releases)
		return out, fmt.Errorf("wrong number of helm releases")
	}

	return fmt.Sprintf("%d", hr), nil
}
