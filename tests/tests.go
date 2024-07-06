package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

var Tests = map[string]func(host string, config []byte) (string, error){
	// Test name: test func.
	"HelmReleases":   HelmReleases,
	"HttpsGet":       HttpsGet,
	"OpenPorts":      OpenPorts,
	"OsRelease":      OsRelease,
	"PodsNotRunning": PodsNotRunning,
}

type test struct {
	name         string
	failedReason string
	err          error
}

func Run(configFile string) (failed int, err error) {
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
		fmt.Printf("--- %s ---\n", host)

		ch := make(chan test, len(testsConfig))

		for testName, testConfig := range testsConfig {
			testFunc, ok := Tests[testName]
			if !ok {
				return 0, fmt.Errorf("no such test: %s", testName)
			}

			go func(host, testName string, testConfig []byte) {
				var t test
				t.name = testName
				t.failedReason, t.err = testFunc(host, testConfig)
				ch <- t
			}(host, testName, testConfig)
		}

		for range testsConfig {
			t := <-ch
			if t.err != nil {
				return 0, fmt.Errorf("run test %s against %s: %v", t.name, host, t.err)
			}
			if t.failedReason != "" {
				failed++
				printFail(t)
			} else {
				printOk(t)
			}
		}
	}

	return failed, nil
}

func printFail(t test) {
	msg := fmt.Sprintf("fail %s", t.name)
	msg += fmt.Sprintf(": %s", t.failedReason)
	fmt.Println(msg)
}

func printOk(t test) {
	msg := fmt.Sprintf("ok   %s", t.name)
	fmt.Println(msg)
}
