package goiris

import (
	"bytes"
	"net/http"
	"net/url"
	"io"
)

type APIClient struct {
	AuthStrategy AuthStrategy
	BaseURL      string
	Client       MyHttpClient
}

func (client *APIClient) DoRequest(builder RequestBuilder) (*http.Response, error) {
	url, err := url.JoinPath(client.BaseURL, builder.URL)
	if err != nil {
		return nil, err
	}

	var body io.Reader = http.NoBody
	if builder.Body != nil {
		body = bytes.NewReader(builder.Body.([]byte))
	}

	req, err := http.NewRequest(builder.Method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range builder.Headers {
		req.Header.Set(k,v)
	}

	client.AuthStrategy.Authenticate(req)
	return client.Client.Do(req)
}
