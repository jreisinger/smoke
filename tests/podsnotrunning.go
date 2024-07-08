package tests

import (
	"encoding/json"
	"fmt"

	"github.com/jreisinger/smoke/helper"
)

type podsNotRunning int

func PodsNotRunning(hostName string, config []byte) (string, bool) {
	var pnr podsNotRunning
	if err := json.Unmarshal(config, &pnr); err != nil {
		return fmt.Sprintf("unmarshal NonRunningPods config: %v", err), true
	}

	cmd := "kubectl get pods --field-selector status.phase!=Running,status.phase!=Succeeded --all-namespaces --no-headers"
	out, err := helper.Ssh(hostName, cmd)
	if err != nil {
		return fmt.Sprintf("ssh %q: %s", cmd, err), false
	}

	podsNotRunning := helper.CountNonEmptyLines(out)
	if podsNotRunning != int(pnr) {
		out := fmt.Sprintf("want count %d, got count %d", pnr, podsNotRunning)
		return out, false
	}

	return fmt.Sprintf("%d", pnr), true
}
