package iris

import (
	"net/http"
)

type AuthStrategy interface {
	Authenticate(req *http.Request)
}

type ApiKeyAuth struct {
	ApiKey string
}

func (a *ApiKeyAuth) Authenticate(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+a.ApiKey)
}
