// Package helper contains functions used to run smoke tests.
package helper

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
)

func StringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CountNonEmptyLines(input []byte) int {
	var count int
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		switch line {
		case "", "\n":
		default:
			count++
		}

	}
	return count
}

func Ssh(host, command string) (stdout []byte, err error) {
	// Load user's SSH config file
	f, err := os.Open(os.ExpandEnv("$HOME/.ssh/config"))
	sshConfig, err := ssh_config.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to load SSH config: %s", err)
	}

	// Get the values from username's SSH config file
	username, err := sshConfig.Get(host, "User")
	if err != nil {
		return nil, fmt.Errorf("failed to get ssh User from config: %s", err)
	}
	identityFile, err := sshConfig.Get(host, "IdentityFile")
	if err != nil {
		return nil, fmt.Errorf("failed to get ssh IndentityFile from config: %s", err)
	}

	if identityFile == "" {
		identityFile = "~/.ssh/id_rsa"
	}

	// Expand tilde to home directory
	if strings.HasPrefix(identityFile, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		identityFile = filepath.Join(dir, identityFile[2:])
	}

	// Load private key for authentication
	keyBytes, err := os.ReadFile(identityFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %s", err)
	}

	key, err := ssh.ParsePrivateKey(keyBytes)
	// Or with password
	// key, err := ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte("password"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	addr := net.JoinHostPort(host, "22")
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %v", addr, err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		return nil, fmt.Errorf("failed to run: %v", err)
	}
	return b.Bytes(), nil
}
