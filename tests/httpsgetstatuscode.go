package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpsGetStatusCode int

func HttpsGetStatusCode(hostName string, config []byte) (string, error) {
	var code httpsGetStatusCode
	if err := json.Unmarshal(config, &code); err != nil {
		return "", fmt.Errorf("unmarshal HttpsGet config: %v", err)
	}

	u := fmt.Sprintf("https://%s", hostName)
	resp, err := http.Get(u)
	msg := fmt.Sprintf("%d, want %d", resp.StatusCode, code)
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != int(code) {
		return msg, fmt.Errorf("wrong status code")
	}
	return fmt.Sprintf("%d", code), nil
}
