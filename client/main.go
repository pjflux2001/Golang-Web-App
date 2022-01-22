package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Plain old URL - http://localhost
	response, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// URL key-value form - http://localhost/url?name=Wallace
	response, err = http.PostForm(
		"http://localhost/url",
		url.Values{"name": {"Wallace"}})
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// http://localhost/body
	// with body in json: {"name": "Wallace"}
	type Ninja struct {
		Name string
	}
	wallace := Ninja{"Wallace"}
	wallaceJson, _ := json.Marshal(wallace)
	response, err = http.Post(
		"http://localhost/body",
		"application/json",
		bytes.NewBuffer(wallaceJson))
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	client := http.Client{}
	request, err := http.NewRequest(
		"GET",
		"http://localhost/body",
		bytes.NewBuffer(wallaceJson))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

}
