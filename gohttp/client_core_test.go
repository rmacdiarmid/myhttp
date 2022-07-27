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
	if finalHeaders.Get("X-Request-Id") != "Abc-123" {
		t.Error("invalid request id received")
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
}
func TestGetRequestBodyNilBody(t *testing.T) {
	//initialization
	client := httpClient{}
	//execution
	body, err := client.getRequestBody("", nil)
	//validation
	if err != nil {
		t.Error("no error expected when passing a nil body")
	}
	if body != nil {
		t.Error("no body expected when a nil passing body")
	}

}
func TestGetRequestBodyWithJson(t *testing.T) {
	//initialization
	//execution
	//validation
}
func TestGetRequestBodyWithXml(t *testing.T) {
	//initialization
	//execution
	//validation
}
func TestGetRequestBodyWithJsonAsDefault(t *testing.T) {
	//initialization
	//execution
	//validation
}
