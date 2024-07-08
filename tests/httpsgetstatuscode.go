package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpsGetStatusCode int

func HttpsGetStatusCode(hostName string, config []byte) (string, bool) {
	var code httpsGetStatusCode
	if err := json.Unmarshal(config, &code); err != nil {
		return fmt.Sprintf("unmarshal HttpsGet config: %v", err), false
	}

	u := fmt.Sprintf("https://%s", hostName)
	resp, err := http.Get(u)
	if err != nil {
		return err.Error(), false
	}
	defer resp.Body.Close()

	if resp.StatusCode != int(code) {
		msg := fmt.Sprintf("%d, want %d", resp.StatusCode, code)
		return msg, false
	}

	return fmt.Sprintf("%d", code), true
}
