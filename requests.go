package ciphertrust

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAll(endpoint string) ([]Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	resp := []Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
