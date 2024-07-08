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
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != int(code) {
		msg := fmt.Sprintf("%d, want %d", resp.StatusCode, code)
		return msg, err
	}
	return fmt.Sprintf("%d", code), nil
}
