package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ninja struct {
	Name string `json:"name"`
}

func url(w http.ResponseWriter, r *http.Request) {
	fmt.Print("url: ")
	name := r.FormValue("name")
	wallace := Ninja{name}
	fmt.Println(wallace)
	wallaceJson, _ := json.Marshal(wallace)
	w.Write(wallaceJson)
}

func body(w http.ResponseWriter, r *http.Request) {
	fmt.Print("body: ")
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var wallace Ninja
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		decoder.Decode(&wallace)
		fmt.Println(wallace)
		wallaceJson, _ := json.Marshal(wallace)
		w.Write(wallaceJson)
	default:
		fmt.Println("Not supported content type", r.Header.Get("Content-Type"))
	}
}

// {"name": "Wallace"}
func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
		// http://localhost/url?name=Wallace
	case "/body":
		body(w, r)
		// http://localhost/body
		// with body in json: {"name": "Wallace"}
	default:
		w.Write([]byte("Hello World"))
		// http://localhost
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
