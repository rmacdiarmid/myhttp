package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rmacdiarmid/myhttp/gohttp"
)

var (
	httpClient = gohttp.New()
)

func main() {
	headers := make(http.Header)

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

}
