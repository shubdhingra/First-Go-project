package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://reqres.in/api/users/"
	var client = http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error...")
	} else {
		fmt.Println("resp", resp)
		data1, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("\n\n\ndata is", string(data1))

	}
}
