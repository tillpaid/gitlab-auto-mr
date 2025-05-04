package httpclient

import (
	"encoding/base64"
	"net/http"
)

type AuthFunc func(req *http.Request)

func BearerAuth(token string) AuthFunc {
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Bearer "+token)
	}
}

func BasicAuth(username, password string) AuthFunc {
	encoded := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Basic "+encoded)
	}
}
