package tests

import (
	"encoding/json"
	"fmt"

	"github.com/jreisinger/smoke/helper"
)

type podsNotRunning struct {
	Count int
}

func PodsNotRunning(hostName string, config []byte) (string, error) {
	var pnr podsNotRunning
	if err := json.Unmarshal(config, &pnr); err != nil {
		return "", fmt.Errorf("unmarshal NonRunningPods config: %v", err)
	}

	cmd := "kubectl get pods --field-selector status.phase!=Running,status.phase!=Succeeded --all-namespaces --no-headers"
	out, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), nil
	}

	podsNotRunning := helper.CountNonEmptyLines(out)
	if podsNotRunning != pnr.Count {
		out := fmt.Sprintf("want %d, got %d", pnr.Count, podsNotRunning)
		return out, nil
	}

	return "", nil
}
