package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//initialization
	client := &httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.builder.headers = commonHeaders

	//execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3")
	}
	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
}
func TestGetRequestBody(t *testing.T) {
	//initialization
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		//execution
		body, err := client.getRequestBody("", nil)
		//validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}
		if body != nil {
			t.Error("no body expected when a nil passing body")
		}
	})

	t.Run("bodyWithJson", func(t *testing.T) {
		//execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))

		//validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})

	/* t.Run("bodyWithXml",func(t *testing.T)) {
	//execution
	//validation
	 }

	t.Run("bodyWithJsonAsDefault",func(t *testing.T)) {
	//execution
	//validation
	} */
}
