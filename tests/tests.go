package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

var Tests = map[string]func(host string, config []byte) (string, error){
	// Test name: test func.
	"HelmReleases": HelmReleases,
	"HttpsGet":     HttpsGet,
	"OpenPorts":    OpenPorts,
	"OsRelease":    OsRelease,
}

func Run(configFile string, verbose bool) (failed int, err error) {
	b, err := os.ReadFile(configFile)
	if err != nil {
		return 0, fmt.Errorf("read config file: %v", err)
	}

	// host -> tests config
	var config map[string]map[string]json.RawMessage

	if err := json.Unmarshal(b, &config); err != nil {
		return 0, fmt.Errorf("unmarshal config file: %v", err)
	}

	for host, testsConfig := range config {
		for testName, testConfig := range testsConfig {
			testFunc, ok := Tests[testName]
			if !ok {
				return 0, fmt.Errorf("no such test: %s", testName)
			}

			faileReason, err := testFunc(host, testConfig)
			if err != nil {
				return 0, fmt.Errorf("run test %s against %s: %v", testName, host, err)
			}

			if faileReason != "" {
				failed++
				msg := fmt.Sprintf("fail %s on %s", testName, host)
				if verbose {
					msg += fmt.Sprintf(": %s", faileReason)
				}
				fmt.Println(msg)
			} else if verbose {
				msg := fmt.Sprintf("ok   %s on %s", testName, host)
				fmt.Println(msg)
			}
		}
	}

	return failed, nil
}
