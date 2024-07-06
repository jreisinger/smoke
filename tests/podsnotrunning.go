package tests

import (
	"encoding/json"
	"fmt"

	"github.com/jreisinger/smoke/helper"
)

type podsNotRunning struct {
	Count int
}

func PodsNotRunning(host string, config []byte) (string, error) {
	var pnr podsNotRunning
	err := json.Unmarshal(config, &pnr)
	if err != nil {
		return "", fmt.Errorf("unmarshal NonRunningPods config: %v", err)
	}

	cmd := "kubectl get pods --field-selector status.phase!=Running --all-namespaces"
	out, err := helper.Ssh(host, cmd)
	if err != nil {
		return "", err
	}

	podsNotRunning := helper.CountNonEmptyLines(out)
	if podsNotRunning != pnr.Count {
		out := fmt.Sprintf("want %d, got %d", pnr.Count, podsNotRunning)
		return out, nil
	}

	return "", nil
}
