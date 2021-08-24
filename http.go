package revoltgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Send http request
func (c Client) Request(method, path string, data []byte) ([]byte, error) {
	reqBody := bytes.NewBuffer(data)

	// Prepare request
	req, err := http.NewRequest(method, API_URL+path, reqBody)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("X-Bot-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.HTTP.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return []byte{}, fmt.Errorf("%s: %s", resp.Status, body)
	}

	return body, nil
}
