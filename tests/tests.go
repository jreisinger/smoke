// Package tests contains smoke tests you can run against a host.
package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

type hostName string
type testName string

type ConfigFile map[hostName]map[testName]json.RawMessage

type TestFunc func(hostName string, config []byte) (msg string, err error)

var Available = map[testName]TestFunc{
	"FilesPresent":       FilesPresent,
	"HelmReleases":       HelmReleases,
	"HttpsGetStatusCode": HttpsGetStatusCode,
	"PortsOpen":          PortsOpen,
	"OsRelease":          OsRelease,
	"PodsNotRunning":     PodsNotRunning,
}

type test struct {
	name string
	msg  string
	err  error
}

func Run(configFilename string) (failed int, err error) {
	b, err := os.ReadFile(configFilename)
	if err != nil {
		return 0, fmt.Errorf("read config file: %v", err)
	}

	var configFile ConfigFile
	if err := json.Unmarshal(b, &configFile); err != nil {
		return 0, fmt.Errorf("unmarshal config file: %v", err)
	}

	for host, tests := range configFile {
		fmt.Printf("--- %s ---\n", host)

		ch := make(chan test, len(tests))

		for testName, testConfig := range tests {
			testFunc, ok := Available[testName]
			if !ok {
				return 0, fmt.Errorf("no such test: %s", testName)
			}

			go func(testName, host string, testConfig []byte) {
				var t test
				t.name = testName
				t.msg, t.err = testFunc(host, testConfig)
				ch <- t
			}(string(testName), string(host), testConfig)
		}

		for range tests {
			t := <-ch
			if t.err != nil {
				failed++
				fmt.Printf("fail %s: %s\n", t.name, t.msg)
			} else {
				fmt.Printf("ok   %s: %s\n", t.name, t.msg)
			}
		}
	}

	return failed, nil
}
