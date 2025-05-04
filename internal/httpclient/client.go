package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	auth       AuthFunc
}

func NewClient(baseUrl string, auth AuthFunc) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseUrl,
		auth:       auth,
	}
}

func (c *Client) DoGetAndDecode(path string, target interface{}) error {
	url := c.baseURL + path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if c.auth != nil {
		c.auth(req)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
