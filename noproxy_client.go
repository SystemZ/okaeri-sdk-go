package oksdk

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type NoProxyClient struct {
	// set your own http client for custom timeouts etc
	HttpClient *http.Client
	// set your own address for API, mainly for debugging
	Endpoint string
	// request API key from https://manager.okaeri.eu or ask friendly Okaeri staff
	ApiKey string
	// enable or disable debugging
	LogEnabled bool
	// pass your own function for handling errors
	LogFunc func(data interface{})
}

func NewNoProxyClient(apiKey string, debug bool) (*NoProxyClient, error) {
	httpClient := &http.Client{}
	return &NoProxyClient{
		HttpClient: httpClient,
		Endpoint:   "https://noproxy-api.okaeri.eu/v1/",
		ApiKey:     apiKey,
		LogEnabled: debug,
		LogFunc: func(data interface{}) {
			log.Printf("%v", data)
		},
	}, nil
}

func (c *NoProxyClient) Get(ctx context.Context, path string) ([]byte, *http.Response, error) {
	return c.Request(ctx, "GET", path, []byte{})
}

func (c *NoProxyClient) Request(ctx context.Context, method string, path string, body []byte) ([]byte, *http.Response, error) {
	reqUrl := c.Endpoint + path
	r, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("fail when creating request for %v: %v", path, err)
	}
	r.WithContext(ctx)
	r.Header.Set("Authorization", "Bearer "+c.ApiKey)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "systemz-ok-sdk-v0")
	res, _ := c.HttpClient.Do(r)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	responseData, err := ioutil.ReadAll(res.Body)
	if err == nil && c.LogEnabled {
		c.LogFunc(fmt.Sprintf("HTTP %v @ %v : %s", res.StatusCode, reqUrl, responseData))
	}
	return responseData, res, nil
}
