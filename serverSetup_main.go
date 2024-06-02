package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest () {
	const myurl = "http://localhost:8000/get"

	res,err := http.Get(myurl)

	if err!=nil {
		panic(err);
	}

	defer res.Body.Close();
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

func PerformPostJsonRequest () {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader (`{
		"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
	}`)

	res,err := http.Post(myurl,"application/json",requestBody)

	if err!=nil {
		panic(err);
	}

	defer res.Body.Close();

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

}

func PerformPostFormRequest () {
	const myurl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}