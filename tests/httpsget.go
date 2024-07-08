package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpsGet struct {
	StatusCode int
}

func HttpsGet(hostName string, config []byte) (string, error) {
	var hg httpsGet
	if err := json.Unmarshal(config, &hg); err != nil {
		return "", fmt.Errorf("unmarshal HttpsGet config: %v", err)
	}

	u := fmt.Sprintf("https://%s", hostName)
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	msg := fmt.Sprintf("%s -> %d (want %d)", u, resp.StatusCode, hg.StatusCode)
	if resp.StatusCode != hg.StatusCode {
		return msg, err
	}
	return msg, nil
}
