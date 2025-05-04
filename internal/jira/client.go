package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    string
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(baseURL, username, password string) *Client {
	return &Client{
		baseURL:    baseURL,
		username:   username,
		password:   password,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetIssueByKey(issueKey string) (*Issue, error) {
	path := fmt.Sprintf("/rest/api/2/issue/%s", url.PathEscape(issueKey))

	var issue Issue
	if err := c.doGetAndDecode(path, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (c *Client) doGetAndDecode(path string, target interface{}) error {
	fullUrl := c.baseURL + path

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.basicAuthHeader())
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
