package main

import (
	"fmt"
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
	for i := 0; i < 1; i++ {
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

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

	/* // Using our custom response
	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}
	fmt.Println(user.FirstName) */

	/* //Using default http.Response

	response.Body.Close()

	fmt.Println(response.StatusCode)

	bytes2, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var user2 User
	if err := json.Unmarshal(bytes, &user); err != nil {
		panic(err)
	}
	fmt.Println(user.FirstName) */
}

func createUser(user User) {

	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	/* bytes, _ := ioutil.ReadAll(response.b)
	fmt.Println(string(bytes)) */

}
