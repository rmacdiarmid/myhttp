package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/rmacdiarmid/myhttp/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//Initialization:
		mock := gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		}
		//Execution:
		endpoints, err := GetEndpoints()
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		//Initialization:
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       {"current_user_url": 123},
		}
		//Execution:
		endpoints, err := GetEndpoints()
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization:
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       {"current_user_url": "https://api.github.com/user"},
		}
		//Execution:
		endpoints, err := GetEndpoints()
	})
}
