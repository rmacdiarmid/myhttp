package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//initialization
	client := &httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	//execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "Abc-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3")
	}
}
