package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jreisinger/smoke/ssh"
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

	out, err := ssh.Ssh(host, "helm ls -A")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	releases := len(lines) - 2 // don't count header line and last empty line
	if releases != hr.Count {
		out := fmt.Sprintf("want %d, got %d", hr.Count, releases)
		return out, nil
	}

	return "", nil
}