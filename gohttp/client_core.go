package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	//Add common headers to the request:
	for header, value := range headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}

	//Add custom headers to the request
	for header, value := range headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}
	return client.Do(request)

}
