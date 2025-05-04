package jira

import "encoding/base64"

func (c *Client) basicAuthHeader() string {
	auth := c.username + ":" + c.password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
