package jira

import (
	"fmt"
	"net/url"

	"github.com/tillpaid/gitlab-auto-mr/internal/httpclient"
)

type Client struct {
	httpClient *httpclient.Client
}

func NewClient(baseURL, username, password string) *Client {
	return &Client{
		httpClient: httpclient.NewClient(baseURL, httpclient.BasicAuth(username, password)),
	}
}

func (c *Client) GetIssueByKey(issueKey string) (*Issue, error) {
	path := fmt.Sprintf("/rest/api/2/issue/%s", url.PathEscape(issueKey))

	var issue Issue
	if err := c.httpClient.DoGetAndDecode(path, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
