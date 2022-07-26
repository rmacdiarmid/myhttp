package main

import (
	"fmt"
	"io/ioutil"

	"github.com/rmacdiarmid/myhttp/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Password Blah Blah")
	// client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

}
