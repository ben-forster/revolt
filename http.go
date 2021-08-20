package revoltgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// Send http request
func (c *Client) Request(method, path, data string) ([]byte, error) {
	var reqBody *bytes.Buffer

	// Check method
	if strings.EqualFold(method, "get") {
		reqBody = bytes.NewBuffer([]byte(data))
	} else {
		reqBody = nil
	}

	// Prepare request
	req, err := http.NewRequest(method, API_URL+path, reqBody)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("X-Bot-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// Send request
	resp, err := client.Do(req)

	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []byte(""), err
	}

	return body, nil
}
