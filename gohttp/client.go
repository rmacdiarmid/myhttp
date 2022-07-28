package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client

	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool

	Headers http.Header
}

func New() HttpClient {
	httpClient := &httpClient{}
	return httpClient

}

type HttpClient interface {
	//Configurations
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(i int)
	DisableTimeouts(disable bool)

	//Http Calls
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

//These are the common headers that we use on every request
func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

//Setting Timeout customization
func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetResponseTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *httpClient) SetMaxIdleConnections(i int) {
	c.maxIdleConnections = i
}

//Disable Timeouts
func (c *httpClient) DisableTimeouts(disable bool) {
	c.disableTimeouts = disable
}

//Basic CRUD Methods setup from HttpClient struct
func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
