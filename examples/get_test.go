package examples

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/rmacdiarmid/myhttp/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	//Tell the HTTP library to mock any further requests from here.
	gohttp.StartMockServer()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})
		//Execution:
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("No endpoints expected at thsi point")
		}
		if err == nil {
			t.Error("an error was expected")
		}
		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message recieved")
		}
	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		//Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})
		//Execution:
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("No endpoints expected at thsi point")
		}
		if err == nil {
			t.Error("an error was expected")
		}
		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message recieved")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		//Execution:
		endpoints, err := GetEndpoints()

		//Validation
		if err == nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))

		}
		if endpoints != nil {
			t.Error("Endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "timeout getting github endpoints" {
			t.Error("invalid current user url")
		}
	})
}
