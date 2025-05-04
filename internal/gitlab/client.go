package gitlab

import (
	"fmt"
	"net/url"

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

func (c *Client) CreateMergeRequest(assigneeId int, projectPath, branch, title, description string) (*MergeRequest, error) {
	path := fmt.Sprintf("/api/v4/projects/%s/merge_requests", url.PathEscape(projectPath))
	body := CreateMergeRequestRequest{
		SourceBranch:       branch,
		TargetBranch:       "master",
		Title:              title,
		Description:        description,
		AssigneeId:         assigneeId,
		Squash:             true,
		RemoveSourceBranch: true,
	}

	var mergeRequest MergeRequest
	if err := c.httpClient.DoPostAndDecode(path, body, &mergeRequest); err != nil {
		return nil, err
	}

	return &mergeRequest, nil
}
