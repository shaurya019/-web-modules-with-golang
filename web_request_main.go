package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")
	res,err := http.Get(url);

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", res)

	defer res.Body.Close()


	dataBytes, err := ioutil.ReadAll(res.Body);

	if err != nil {
		panic(err)
	}

	content := string(dataBytes);
	fmt.Println(content)



}