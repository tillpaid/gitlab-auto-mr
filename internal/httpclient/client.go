package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	req.Header.Set("Accept", "application/json")
	if c.auth != nil {
		c.auth(req)
	}

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

func (c *Client) DoPostAndDecode(path string, body interface{}, target interface{}) error {
	url := c.baseURL + path

	var buf io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.auth != nil {
		c.auth(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return ParseAPIError(resp.StatusCode, responseBody)
	}

	if target != nil {
		return json.Unmarshal(responseBody, target)
	}

	return nil
}
