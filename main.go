package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World, Golang!</h1>\n")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello</h1>"+" /\n"+r.UserAgent()+"\n"+r.Referer()+"\n")
	case "/ninja":
		fmt.Fprint(w, "Ninja here")
	default:
		fmt.Fprint(w, "Err")
	}
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not timeout")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.HandleFunc("/timeout", timeout)
	// server := http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      nil,
	// 	ReadTimeout:  1000,
	// 	WriteTimeout: 1000,
	// }
	http.ListenAndServe(":8080", nil)
	// server.ListenAndServe()
}
