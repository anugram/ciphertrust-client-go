package ciphertrust

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAll(endpoint string) ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.CipherTrustURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	resp := []User{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
