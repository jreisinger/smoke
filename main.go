// Smoke is a simple tool for running basic sanity tests against one or more hosts.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jreisinger/smoke/tests"
)

var defaultConfigFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "smoke: get user home dir: %v\n", err)
		os.Exit(1)
	}
	defaultConfigFile = filepath.Join(homeDir, ".smoke.json")
}

func main() {
	c := flag.String("c", defaultConfigFile, "config file")
	flag.Parse()

	failed, err := tests.Run(*c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "smoke: run tests: %v\n", err)
		os.Exit(1)
	}
	os.Exit(failed)
}
