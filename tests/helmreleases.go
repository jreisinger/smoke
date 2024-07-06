package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/helper"
)

type helmReleases struct {
	Count int
}

func HelmReleases(host string, config []byte) (string, error) {
	var hr helmReleases
	err := json.Unmarshal(config, &hr)
	if err != nil {
		return "", fmt.Errorf("unmarshal HelmReleases config: %v", err)
	}

	cmd := "helm ls -A"
	out, err := helper.Ssh(host, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), nil
	}

	lines := strings.Split(string(out), "\n")
	releases := len(lines) - 2 // don't count header line and last empty line
	if releases != hr.Count {
		out := fmt.Sprintf("want %d, got %d", hr.Count, releases)
		return out, nil
	}

	return "", nil
}
