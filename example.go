package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rmacdiarmid/myhttp/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			getUrls()
		}()
	}
	time.Sleep(20 * time.Second)
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

func createUser(user User) {

	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

}
