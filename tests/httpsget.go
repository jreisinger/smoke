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
	err := json.Unmarshal(config, &hg)
	if err != nil {
		return "", fmt.Errorf("unmarshal HttpsGet config: %v", err)
	}

	u := fmt.Sprintf("https://%s", hostName)
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != hg.StatusCode {
		return fmt.Sprintf("want %d, got %d", hg.StatusCode, resp.StatusCode), nil
	}
	return "", nil
}
