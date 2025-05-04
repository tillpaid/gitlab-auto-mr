package gitlab

import (
	"fmt"

	"github.com/tillpaid/gitlab-auto-mr/internal/httpclient"
)

type Client struct {
	httpClient *httpclient.Client
}

func NewClient(baseURL string, token string) *Client {
	return &Client{
		httpClient: httpclient.NewClient(baseURL, httpclient.BearerAuth(token)),
	}
}

func (c *Client) GetCurrentUser() (*User, error) {
	path := fmt.Sprintf("/api/v4/user")

	var user User
	if err := c.httpClient.DoGetAndDecode(path, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
