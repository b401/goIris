package goiris

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MyHttpClient struct {
	client *http.Client
}

func (mhc *MyHttpClient) Do(req *http.Request) (*http.Response, error) {
	return mhc.client.Do(req)
}

func NewMyHttpClient() *MyHttpClient {
	return &MyHttpClient{client: &http.Client{}}
}

type ClientConfig struct {
	Timeout   time.Duration
	IgnoreTLS bool
}

func NewConfiguredHttpClient(config ClientConfig) *MyHttpClient {
	client := &http.Client{
		Timeout: config.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.IgnoreTLS},
		},
	}

	return &MyHttpClient{client: client}
}
